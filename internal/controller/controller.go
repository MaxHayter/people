package controller

import (
	"github.com/MaxHayter/people/internal/db"
	"github.com/MaxHayter/people/internal/service"
	api "github.com/MaxHayter/people/people"
)

type Controller struct {
	db      *db.StorageFactory
	service *service.Service
	api.UnimplementedPeopleServiceServer
}

func NewController(db *db.StorageFactory, service *service.Service) *Controller {
	return &Controller{
		db:      db,
		service: service,
	}
}
