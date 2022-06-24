package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/RipudamanKaushikDal/MultilangAPI/models"
)

type Task struct {
	TaskURL string `json:"task_status"`
}

type APIResult struct {
	TaskId     string                   `json:"task_id"`
	TaskStatus string                   `json:"task_status"`
	TaskResult []map[string]interface{} `json:"task_result"`
}

type Result struct {
	InvestorID uint                     `json:"investorId"`
	StockInfo  []map[string]interface{} `json:"stockInfo"`
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

func GetTaskId(stockModels []models.Stock) (*http.Response, error) {
	var stockList []string
	for _, stocks := range stockModels {
		stockList = append(stockList, stocks.Symbol)
	}

	stockMap := make(map[string][]string)
	stockMap["stocks"] = stockList

	postdata, _ := json.Marshal(stockMap)

	response, err := http.Post("http://localhost:5004/tasks", "application/json", bytes.NewBuffer(postdata))
	if err != nil {
		panic(err.Error())
	}

	return response, err
}

func GetStockInfo(stockModels []models.Stock) Result {

	var result Result

	result.InvestorID = stockModels[0].InvestorID
	response, err := GetTaskId(stockModels)

	var task Task

	err = getJson(response, &task)
	if err != nil {
		panic(err.Error())
	}

	var apiResults APIResult
	fetchResults := func() {
		response, _ = http.Get("http://localhost:5004" + task.TaskURL)
		err = getJson(response, &apiResults)
		if err != nil {
			panic(err.Error())
		}
	}

	clear := setInterval(fetchResults, 1000)

	for {

		time.Sleep(500)

		if apiResults.TaskStatus == "SUCCESS" {
			clear <- true
			result.StockInfo = apiResults.TaskResult
			return result
		}

	}

}
