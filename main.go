package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"phoenix/route"
)

func main() {
	f, _ := os.Create("log/log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	server := gin.Default()

	router := route.NewRouter(server)
	router.Routes()

	server.Run("127.0.0.1:8080")
}
