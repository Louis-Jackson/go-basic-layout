package main

import (
	"fmt"
	"go-basic-layout/config"
	"go-basic-layout/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()

	r := gin.Default()

	r.GET("/hello", api.HelloHandler)

	r.POST("/hello", api.HelloHandler)

	fmt.Printf("Server is starting on port %d...\n", config.Server.Port)
	r.Run(fmt.Sprintf(":%d", config.Server.Port))
}
