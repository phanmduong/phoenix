package router

import (
	"strings"
	"reflect"
	"github.com/gin-gonic/gin"
	"nimbus/core"
)

type Router struct {
}

func (router *Router) GET(path string, handler string) *Router {
	context := core.GetContext()
	server := context.Server

	controller, method := handle(handler)

	server.GET(path, func(context *gin.Context) {
		inputs := make([]reflect.Value, 1)

		inputs[0] = reflect.ValueOf(context)
		controller.MethodByName(method).Call(inputs)
	})

	return router
}

func (router *Router) POST(path string, handler string) *Router {
	context := core.GetContext()
	server := context.Server

	controller, method := handle(handler)

	server.POST(path, func(context *gin.Context) {
		inputs := make([]reflect.Value, 1)

		inputs[0] = reflect.ValueOf(context)
		controller.MethodByName(method).Call(inputs)
	})

	return router
}

func (router *Router) PUT(path string, handler string) *Router {
	context := core.GetContext()
	server := context.Server

	controller, method := handle(handler)

	server.PUT(path, func(context *gin.Context) {
		inputs := make([]reflect.Value, 1)

		inputs[0] = reflect.ValueOf(context)
		controller.MethodByName(method).Call(inputs)
	})

	return router
}

func (router *Router) DELETE(path string, handler string) *Router {
	context := core.GetContext()
	server := context.Server

	controller, method := handle(handler)

	server.DELETE(path, func(context *gin.Context) {
		inputs := make([]reflect.Value, 1)

		inputs[0] = reflect.ValueOf(context)
		controller.MethodByName(method).Call(inputs)
	})

	return router
}

func handle(handler string) (reflect.Value, string) {
	context := core.GetContext()

	registryManager := context.RegistryManager

	arr := strings.Split(handler, "@")

	controllerName, method := arr[0], arr[1]

	var typeController reflect.Type
	typeController = registryManager.GetControllerRegistry(controllerName)

	controller := reflect.New(typeController)

	return controller, method
}
