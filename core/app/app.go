package app

import (
	"nimbus/registry"
	userRouter "nimbus/router"
	"nimbus/core/router"
	"nimbus/core"
	"nimbus/core/database"
)

type App struct {
	context *core.Context
	Router  router.Router
	DB *database.DatabaseFacade
}

func NewApp() *App {
	app := &App{
		context: core.GetContext(),
		Router:  router.Router{},
		DB: database.NewDatabase(),
	}
	return app
}

func (app *App) Init() {
	app.context.RegistryManager.RegisterControllerRegistry(registry.GetControllerRegistry())
	userRouter.RegisterRoutes(app.Router)
}

func (app *App) Run() {
	server := app.context.Server
	server.Run("127.0.0.1:8080")
}
