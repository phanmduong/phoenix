package main

import (
	app2 "nimbus/core/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"nimbus/core"
)

func main() {
	//f, _ := os.Create("log/log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	app := app2.NewApp()

	app.Init()

	server := core.GetContext().Server

	server.Static("/assets", "./public/assets")
	server.LoadHTMLGlob("views/*")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	app.Run()
}
