package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControler struct {
}

func (controler *UserControler) GetUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"10": 11})
}

//func getData()  {
//
//}
