package service

import (
	"context"

	"github.com/MaxHayter/people/internal/service/entity"
)

type Storage interface {
	CreatePerson(ctx context.Context, r *entity.Person) error
	GetPerson(ctx context.Context, login string) (*entity.Person, error)
	ExistPerson(ctx context.Context, login string) (bool, error)
}
