package models

import "gorm.io/gorm"

type Investor struct {
	gorm.Model
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Stocks []Stock `json:"stocks"`
}

type Stock struct {
	gorm.Model
	InvestorID uint   `json:"investorID"`
	Symbol     string `json:"symbol"`
}
