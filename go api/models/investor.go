package models

import "time"

type Investor struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" binding:"required"`
	Stocks    []string  `json:"stocks" binding:"required"`
	UpdatedAt time.Time `json:"timestamp"`
}
