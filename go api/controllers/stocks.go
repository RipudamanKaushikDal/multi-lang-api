package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/RipudamanKaushikDal/MultilangAPI/models"
	"github.com/gin-gonic/gin"
)

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

func getJson(response *http.Response, structure interface{}) error {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	return json.Unmarshal(body, &structure)
}

func FindInvestors(ctx *gin.Context) {
	var Investors []models.Investor
	results := models.DB.Find(&Investors)
	if results.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": results.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": results})
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

func GetStocks(ctx *gin.Context) {
	stocks, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Stocks:", stocks)
	response, err := http.Post("http://localhost:5004/tasks", "application/json", bytes.NewBuffer(stocks))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var task Task
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &task)
	fmt.Println("TaskId:", string(body))

	if err != nil {
		panic(err.Error())
	}
	// get response from json object

	fmt.Printf("%+v\n", task)

	var results Result
	fetchResults := func() {
		response, _ = http.Get("http://localhost:5004" + task.TaskURL)
		body, _ = ioutil.ReadAll(response.Body)
		err = json.Unmarshal(body, &results)
		if err != nil {
			panic(err.Error())
		}
	}

	clear := setInterval(fetchResults, 1000)

	for {

		fmt.Printf("%+v\n", results)

		if results.TaskStatus == "SUCCESS" {
			clear <- true
			ctx.JSON(http.StatusOK, gin.H{"results": results.TaskResult})
			return
		}
	}

}
