package controllers

import (
	"net/http"
	"sync"

	"github.com/RipudamanKaushikDal/MultilangAPI/models"
	"github.com/gin-gonic/gin"
)

func CreateInvestor(ctx *gin.Context) {
	var Investor models.Investor

	if err := ctx.ShouldBindJSON(&Investor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	investor := models.Investor{Name: Investor.Name, Stocks: Investor.Stocks}
	models.DB.Create(&investor)
	ctx.JSON(http.StatusOK, gin.H{"data": investor})
}

func GetAllInvestors(ctx *gin.Context) {
	investors := models.FindAllInvestors()
	ctx.JSON(http.StatusOK, gin.H{"data": investors})
}

func GetAllStocks(ctx *gin.Context) {
	var wg sync.WaitGroup
	var results []Result
	investors := models.FindAllInvestors()

	fetchInvestorStocks := func(stocks []models.Stock) {
		defer wg.Done()
		results = append(results, GetStockInfo(stocks))
	}

	for _, investor := range investors {
		wg.Add(1)
		go fetchInvestorStocks(investor.Stocks)
	}

	wg.Wait()
	ctx.JSON(http.StatusOK, gin.H{"data": results})
}

func GetStockById(ctx *gin.Context) {
	id := ctx.Param("id")
	investor := models.FindInvestorById(id)
	result := GetStockInfo(investor.Stocks)
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
