package route

import (
	"github.com/gin-gonic/gin"
	"phoenix/controller"
)

type UserRoute struct {
	router         *gin.RouterGroup
	userController controller.UserControler
}

func NewUserRoute(router *gin.RouterGroup) UserRoute {
	userRoute := UserRoute{}
	userRoute.userController = controller.UserControler{}
	userRoute.router = router
	return userRoute
}

func (userRoute *UserRoute) Run() {
	router := userRoute.router
	userController := userRoute.userController
	router.GET("info", userController.GetUser)
}
