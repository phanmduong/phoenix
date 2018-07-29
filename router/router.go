package router

import "nimbus/core/router"

func RegisterRoutes(router *router.Router) {
	router.GET("/user", "UserController@GetUser")
}
