package models

import "time"

type Payment struct {
	PaymentID   uint `gorm:"PrimaryKey"`
	OrderID   	uint
	Order     	Order
	UserID      uint
	User        User
	Type 		string  `gorm:"type:varchar(100);column: type"`
	Amount 		float64 `gorm:"column: amount"`
	DatePayment time.Time
	Status	    string `gorm:"type:varchar(100);column: status"`
	CreatedAt 	time.Time
}
