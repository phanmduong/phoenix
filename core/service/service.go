package service

import (
	"nimbus/core/router"
	"nimbus/core/database"
)

type Service struct {
	Router    *router.Router
	DB *database.DatabaseFacade
}

func NewService() *Service {
	return &Service{
		Router:  &router.Router{},
		DB:      database.NewDatabase(),
	}
}