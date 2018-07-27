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

	//tokenString, error := jwt.CreateJWT("dấdsadasdaddas")
	//
	//fmt.Println(tokenString)
	//fmt.Println(error)

	//tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiZOG6pWRzYWRhc2RhZGRhcyJ9.oQ8tsS89AthRJToV_xw0_9RUW8wDeDq9jVZrm4U_jAM"
	//data, error := jwt.ParseJWT(tokenString)
	//fmt.Println(data)
	//fmt.Println(error)

	app.Run()
}
