package service

import (
	"context"
	"errors"

	"github.com/MaxHayter/people/internal/service/entity"
)

func (s *Service) Registrate(ctx context.Context, r *entity.Person) error {
	exist, err := s.db.ExistPerson(ctx, r.Login)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("person with this login already exists")
	}

	r, err = s.password.CheckAndHash(ctx, r)
	if err != nil {
		return err
	}

	return s.db.CreatePerson(ctx, r)
}

func (s *Service) Login(ctx context.Context, r *entity.Person) (bool, error) {
	person, err := s.db.GetPerson(ctx, r.Login)
	if err != nil {
		return false, err
	}

	person.Password = r.Password
	return s.password.Compare(ctx, person)
}
