package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Investors []Investor

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Investor{}, &Stock{})

	DB = database
}

func FindAllInvestors() Investors {
	var investors Investors
	dbSearch := DB.Preload("Stocks").Find(&investors)

	if err := dbSearch.Error; err != nil {
		panic(err)
	}
	return investors
}

func FindInvestorById(id string) Investor {
	var investor Investor
	dbSearch := DB.Preload("Stocks").First(&investor, id)
	if err := dbSearch.Error; err != nil {
		panic(err)
	}
	return investor
}
