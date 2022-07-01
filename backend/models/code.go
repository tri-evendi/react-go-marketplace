package models

import "time"

type Code struct {
	ID        	 uint   `gorm:"primary_key;NOT NULL AUTO_INCREMENT"`
	Code     	 string `gorm:"UNIQUE; NOT NULL"`
	ProductID    uint
	Product      Product
	DateExpired	 time.Time
	IsAvailable	 bool `gorm:"column: is_available"`
}
