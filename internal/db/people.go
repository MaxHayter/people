package db

import (
	"context"
	"errors"

	"github.com/MaxHayter/people/internal/db/model"
	"github.com/MaxHayter/people/internal/service/entity"
	"github.com/MaxHayter/people/logger"
)

func (c Repository) CreatePerson(ctx context.Context, r *entity.Person) error {
	log := logger.GetLogger(ctx)

	person := &model.Person{
		Login:    r.Login,
		Password: r.HashPassword,
	}

	_, err := c.db.ModelContext(ctx, person).
		Insert()
	if err != nil {
		log.Println(err.Error())
		return errors.New("unable to create person")

	}
	return nil
}

func (c Repository) GetPerson(ctx context.Context, login string) (*entity.Person, error) {
	log := logger.GetLogger(ctx)

	person := &model.Person{}

	err := c.db.ModelContext(ctx, person).
		Where(model.Columns.Person.Login+" = ?", login).
		First()
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("unable to get person")

	}
	return &entity.Person{
		Login:        person.Login,
		HashPassword: person.Password,
	}, nil
}

func (c Repository) ExistPerson(ctx context.Context, login string) (bool, error) {
	log := logger.GetLogger(ctx)

	exist, err := c.db.ModelContext(ctx, (*model.Person)(nil)).
		Where(model.Columns.Person.Login+" = ?", login).
		Exists()
	if err != nil {
		log.Println(err.Error())
		return false, errors.New("unable to chech exist person")

	}
	return exist, err
}
