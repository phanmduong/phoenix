package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"phoenix/route"
	"net/http"
)

func main() {
	f, _ := os.Create("log/log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	server := gin.Default()

	server.Static("/assets", "./public/assets")
	server.LoadHTMLGlob("views/*")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router := route.NewRouter(server)
	router.Routes()

	server.Run("127.0.0.1:8080")
}
