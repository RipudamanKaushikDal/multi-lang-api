package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RipudamanKaushikDal/MultilangAPI/models"

	"github.com/RipudamanKaushikDal/MultilangAPI/controllers"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hi There!"})
	})

	router.GET("/investors", controllers.FindInvestors)
	router.POST("/investors", controllers.CreateInvestor)
	router.POST("/stocks", controllers.GetStocks)

	router.Run()
}
