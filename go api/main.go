package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hi There!"})
	})

	router.Run()
}
