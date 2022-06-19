package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/RipudamanKaushikDal/MultilangAPI/models"
	"github.com/gin-gonic/gin"
)

type DatabaseQuery struct {
	Id     uint
	Stocks []string
}

type Task struct {
	TaskURL string `json:"task_status"`
}

type Result struct {
	TaskId     string                   `json:"task_id"`
	TaskStatus string                   `json:"task_status"`
	TaskResult []map[string]interface{} `json:"task_result"`
}

func setInterval(task func(), duration int) chan bool {
	interval := time.Duration(duration) * time.Millisecond

	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				go task()

			case <-clear:
				ticker.Stop()
				return
			}
		}
	}()

	return clear
}

func getJson(response *http.Response, structureReference interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	return json.Unmarshal(body, structureReference)
}

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

func GetAllStocks(ctx *gin.Context) {
	var investor []models.Investor
	dbSearch := models.DB.Preload("Stocks").Find(&investor)

	if err := dbSearch.Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": investor})
}

func GetStocks(ctx *gin.Context) {

	var investor []byte
	response, err := http.Post("http://localhost:5004/tasks", "application/json", bytes.NewBuffer(investor))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task Task

	err = getJson(response, &task)
	if err != nil {
		panic(err.Error())
	}

	var results Result
	fetchResults := func() {
		response, _ = http.Get("http://localhost:5004" + task.TaskURL)
		err = getJson(response, &results)
		if err != nil {
			panic(err.Error())
		}
	}

	clear := setInterval(fetchResults, 1000)

	for {

		time.Sleep(1000)

		if results.TaskStatus == "SUCCESS" {
			clear <- true
			ctx.JSON(http.StatusOK, gin.H{"results": results.TaskResult})
			return
		}
	}

}
