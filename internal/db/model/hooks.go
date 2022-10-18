package model

import (
	"context"

	"github.com/go-pg/pg/v10/orm"
)

// Person hook
var _ orm.BeforeInsertHook = (*Person)(nil)
var _ orm.BeforeUpdateHook = (*Person)(nil)

func (model *Person) BeforeInsert(ctx context.Context) (context.Context, error) {
	model.ID = GenStringUUID()

	return ctx, nil
}

func (model *Person) BeforeUpdate(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
