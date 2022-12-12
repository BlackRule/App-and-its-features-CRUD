package main

import (
	"github.com/gin-gonic/gin"
	"ichernovalov/app_compare/http"
	"os"
)

const defaultPort = ":80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	server := gin.Default()
	server.GET("/playground", http.PlaygroundHandler())
	server.POST("/query", http.GraphQLHandler())
	server.Static("/", "./static")
	server.Run(port)
}
