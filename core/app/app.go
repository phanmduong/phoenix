package app

import (
	"nimbus/registry"
	userRouter "nimbus/router"
	"nimbus/core"
	"nimbus/core/service"
)

type App struct {
	service *service.Service
	context *core.Context
}

func NewApp() *App {
	app := &App{
		context: core.GetContext(),
		service: service.NewService(),
	}
	return app
}

func (app *App) Init() {
	app.context.RegistryManager.RegisterControllerRegistry(registry.GetControllerRegistry())
	userRouter.RegisterRoutes(app.service.Router)
}

func (app *App) Run() {
	server := app.context.Server
	server.Run("127.0.0.1:8080")
}
