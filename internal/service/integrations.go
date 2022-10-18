package service

import (
	"context"

	"github.com/MaxHayter/people/internal/service/entity"
)

type PasswordService interface {
	CheckAndHash(ctx context.Context, r *entity.Person) (*entity.Person, error)
	Compare(ctx context.Context, r *entity.Person) (bool, error)
}
