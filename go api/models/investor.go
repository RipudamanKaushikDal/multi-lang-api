package models

import "gorm.io/gorm"

type Investor struct {
	gorm.Model
	Name   string  `json:"name" binding:"required"`
	Stocks []Stock `json:"stocks" binding:"required"`
}

type Stock struct {
	gorm.Model
	InvestorID uint
	Symbol     string
	Price      string
	Change     string
}
