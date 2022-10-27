package models

import (
	"time"
)

type Order struct {
	Order_ID      uint   `gorm:"primaryKey"`
	Customer_name string `unique;json:"customerName"`
	Items         []Item `json:"items"`
	Ordered_at    time.Time
}
