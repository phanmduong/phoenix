package route

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter(server *gin.Engine) Router {
	rootRouter := Router{}
	rootRouter.router = server
	return rootRouter
}

func (r *Router) Routes() {
	r.userRoutes()
}

func (r *Router) userRoutes() {
	userRoute := NewUserRoute(r.router.Group("/user"))
	userRoute.Run()
}
