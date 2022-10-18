package configs

import (
	"os"

	"github.com/go-pg/pg/v10"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DbConfig    *pg.Options  `yaml:"db_config"`
	Server      *Server      `yaml:"server"`
	Swagger     Swagger      `yaml:"swagger"`
	GrpcClients *GrpcClients `yaml:"grpc_clients"`
}

type GrpcClients struct {
	PasswordEndpointGRPC string `yaml:"password_endpoint_grpc"`
}

type Swagger struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

type Server struct {
	GrpcAddress    string `yaml:"grpc_address"`
	GatewayAddress string `yaml:"gateway_address"`
	BasePath       string `yaml:"base_path"`
	MaxGRPCMsgSize int    `yaml:"max_grpc_msg_size"`
}

func Configure(fileName string) (*Config, error) {
	var cnf *Config
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cnf)
	if err != nil {
		return nil, err
	}

	return cnf, nil
}
