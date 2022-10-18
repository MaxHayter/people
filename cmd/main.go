package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/MaxHayter/password/password"
	"github.com/go-chi/chi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/MaxHayter/people/configs"
	"github.com/MaxHayter/people/internal/controller"
	"github.com/MaxHayter/people/internal/db"
	passwordIntegration "github.com/MaxHayter/people/internal/integrations/password"
	"github.com/MaxHayter/people/internal/service"
	"github.com/MaxHayter/people/logger"
	api "github.com/MaxHayter/people/people"
)

const (
	configPath = "config.yaml"
)

func main() {
	configFile := flag.String("c", configPath, "specify path to a config.yaml")
	flag.Parse()

	log := logger.DefaultLogger
	ctx := logger.WithLogger(context.Background(), log)

	config, err := configs.Configure(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	// create dbConnection connection
	dbConnection, err := db.NewDBConnection(logger.WithLogger(ctx, log), config.DbConfig)
	if err != nil {
		log.Fatal("unable to connect to dbConnection")
	}
	defer func() {
		//nolint:govet shadow declaration is intentional
		err := dbConnection.Close()
		if err != nil {
			log.Println("unable to close db connection")
		}
	}()

	// password grpc client
	passwordConn, err := grpc.DialContext(ctx, config.GrpcClients.PasswordEndpointGRPC,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	)
	if err != nil {
		log.Fatal("unable to connect to PasswordService via grpc")
	}
	defer func() {
		//nolint:govet shadow declaration is intentional
		err := passwordConn.Close()
		if err != nil {
			log.Println("unable to close PasswordService grpc connection")
		}
	}()
	passwordServiceClient := password.NewPasswordServiceClient(passwordConn)

	listen, err := net.Listen("tcp", config.Server.GrpcAddress)
	if err != nil {
		log.Fatal("failed to listen grpc port")
	}

	storage := db.NewStorageFactory(dbConnection)
	peopleController := controller.NewController(storage,
		service.NewService(storage.NewRepository(), passwordIntegration.NewClient(passwordServiceClient)))

	var serverOptions []grpc.ServerOption

	grpcServer := grpc.NewServer(serverOptions...)

	api.RegisterPeopleServiceServer(grpcServer, peopleController)

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption("*", &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				Multiline:       false,
				Indent:          "",
				AllowPartial:    false,
				UseProtoNames:   true,
				UseEnumNumbers:  false,
				EmitUnpopulated: false,
				Resolver:        nil,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				AllowPartial:   true,
				DiscardUnknown: true,
				Resolver:       nil,
			},
		}),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = api.RegisterPeopleServiceHandlerFromEndpoint(context.Background(), mux, config.Server.GrpcAddress, opts)
	if err != nil {
		log.Fatal("register gateway")
	}

	router := chi.NewRouter()
	router.Route(config.Server.BasePath, func(r chi.Router) {
		r.Mount("/", http.StripPrefix(config.Server.BasePath, mux))
	})

	swaggerFullPath := path.Join(config.Server.BasePath, config.Swagger.Url)
	swaggerFullPathPrefix := swaggerFullPath + "/"
	fs := http.FileServer(http.Dir(config.Swagger.Path))
	router.Mount(swaggerFullPathPrefix, http.StripPrefix(swaggerFullPathPrefix, fs))
	router.Get(swaggerFullPath, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, swaggerFullPathPrefix, http.StatusMovedPermanently)
	})

	srv := &http.Server{
		Addr:     config.Server.GatewayAddress,
		Handler:  router,
		ErrorLog: log,
	}

	signalListener := make(chan os.Signal, 1)
	defer close(signalListener)

	signal.Notify(signalListener,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		if err = grpcServer.Serve(listen); err != nil {
			log.Println("failed to listen grpc port")
		}
	}()

	defer func() {
		grpcServer.GracefulStop()
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	defer func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("failed on shutdown server")
		}
	}()

	stop := <-signalListener
	log.Println("Received", stop)
	log.Println("Waiting for all jobs to stop")
}
