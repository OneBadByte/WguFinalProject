package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})

	})
	router.Run()
}
