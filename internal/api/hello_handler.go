package api

import (
	"go-basic-layout/internal/hello"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {

	var requestData struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	greeting := hello.SayHelloTo(requestData.Name)

	c.JSON(200, gin.H{
		"message": greeting,
	})
}
