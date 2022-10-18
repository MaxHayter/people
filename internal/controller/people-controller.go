package controller

import (
	"context"
	"errors"

	"github.com/MaxHayter/people/internal/service/entity"
	"github.com/MaxHayter/people/logger"
	api "github.com/MaxHayter/people/people"
	"github.com/golang/protobuf/ptypes/empty"
)

func (c *Controller) Registrate(ctx context.Context, r *api.Request) (*empty.Empty, error) {
	log := logger.GetLogger(ctx)

	// Если в логике больше 1 обращения к бд, то начинаем транзакцию
	// tx, err := c.db.NewTransaction(ctx)
	// if err != nil {
	// 	return nil, errors.New("unable to start transaction")
	// }
	// c.service = c.service.WithStorage(tx)
	err := c.service.Registrate(ctx, personToEntity(r))
	if err != nil {
		// tx.Rollback(ctx)
		log.Println(err)
		return nil, errors.New("unable to registrate person")
	}

	// err = tx.Commit(ctx)
	// if err != nil {
	// 	return nil, errors.New("unable to commit")
	// }

	return &empty.Empty{}, nil
}

func personToEntity(r *api.Request) *entity.Person {
	return &entity.Person{
		Login:    r.GetLogin(),
		Password: r.GetPassword(),
	}
}

func (c *Controller) Login(ctx context.Context, r *api.Request) (*api.Result, error) {
	log := logger.GetLogger(ctx)

	res, err := c.service.Login(ctx, personToEntity(r))
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to login person")
	}

	return &api.Result{Result: res}, nil
}
