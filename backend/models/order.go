package models

import "time"

type Order struct {
	OrderID     uint `gorm:"PrimaryKey"`
	ProductID   uint
	Product     Product
	UserID      uint
	User        User
	Quantity 	int64   `gorm:"default:1;column: quantity"`
	Amount 		float64 `gorm:"default:0;column: amount"`
	DateOrdered time.Time
	Code 		[]Code `gorm:"many2many:coderables;"`
	Status	    string `gorm:"type:varchar(100);default:pending;column: status"`
	CreatedAt 	time.Time
}
