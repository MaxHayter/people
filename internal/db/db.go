package db

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"

	"github.com/MaxHayter/people/logger"
)

func logQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	log := logger.GetLogger(ctx)
	query, err := q.FormattedQuery()
	if err != nil {
		log.Println(err)
		return ctx, nil
	}
	log.Println("query:", string(query))
	return ctx, nil
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return logQuery(ctx, q)
}

func (d dbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	return nil
}

func NewDBConnection(ctx context.Context, config *pg.Options) (*pg.DB, error) {
	db := pg.Connect(config)

	db.AddQueryHook(dbLogger{})

	_, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		return nil, errors.New("cannot ping Postgres")
	}

	return db, nil
}
