package main

type Investor struct {
	ID     uint8    `json:"id" gorm:"primary_key"`
	Name   string   `json:"name"`
	Stocks []string `json:"stocks"`
}
