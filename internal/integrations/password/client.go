package filter

import (
	"context"

	api "github.com/MaxHayter/password/password"

	"github.com/MaxHayter/people/internal/service/entity"
)

type ClientWrapper struct {
	client api.PasswordServiceClient
}

func NewClient(client api.PasswordServiceClient) *ClientWrapper {
	return &ClientWrapper{client: client}
}

func (c *ClientWrapper) CheckAndHash(ctx context.Context, r *entity.Person) (*entity.Person, error) {
	hash, err := c.client.CheckAndHash(ctx, &api.Request{Request: r.Password})
	if err != nil {
		return nil, err
	}

	r.HashPassword = hash.GetResult()
	return r, nil
}

func (c *ClientWrapper) Compare(ctx context.Context, r *entity.Person) (bool, error) {
	ok, err := c.client.Compare(ctx, &api.CompareRequest{
		Password: r.Password,
		Hash:     r.HashPassword,
	})
	if err != nil {
		return false, err
	}

	return ok.GetOk(), nil
}
