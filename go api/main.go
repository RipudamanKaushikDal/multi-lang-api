package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RipudamanKaushikDal/MultilangAPI/models"

	"github.com/RipudamanKaushikDal/MultilangAPI/controllers"

	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:3000"}

	models.ConnectDatabase()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hi There!"})
	})

	router.POST("/investors", controllers.CreateInvestor)
	router.GET("/investors", controllers.GetAllInvestors)
	router.GET("/stocks", controllers.GetAllStocks)
	router.GET("/stocks/:id", controllers.GetStockById)

	router.Use(cors.New(config))
	router.Run()
}
