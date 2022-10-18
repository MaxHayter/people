package db

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	"github.com/MaxHayter/people/logger"
)

type StorageFactory struct {
	db *pg.DB
}

func NewStorageFactory(db *pg.DB) *StorageFactory {
	return &StorageFactory{
		db: db,
	}
}

func (f *StorageFactory) NewRepository() Repository {
	return Repository{
		db: f.db,
	}
}

func (f *StorageFactory) NewTransaction(ctx context.Context) (RepositoryTx, error) {
	log := logger.GetLogger(ctx)

	tx, err := f.db.WithContext(ctx).Begin()
	if err != nil {
		log.Println("unable to begin transaction")
		return RepositoryTx{}, errors.New("unable to begin transaction")
	}

	return RepositoryTx{
		tx: tx,
		Repository: Repository{
			db: tx,
		},
	}, nil
}

type Repository struct {
	db orm.DB
}

type RepositoryTx struct {
	tx *pg.Tx
	Repository
}

func (c RepositoryTx) Commit(ctx context.Context) error {
	log := logger.GetLogger(ctx)

	err := c.tx.Commit()
	if err != nil {
		log.Println("unable to commit transaction")
		return errors.New("unable to commit transaction")
	}

	return nil
}

func (c RepositoryTx) Rollback(ctx context.Context) {
	log := logger.GetLogger(ctx)

	err := c.tx.Rollback()
	if err != nil {
		log.Println("unable to rollback transaction")
	}
}
