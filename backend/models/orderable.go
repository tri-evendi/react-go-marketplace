package models

import "time"

type Orderable struct {
	OrderID   	uint
	Order     	Order
	ProductID	uint
	Product     Product
	Quantity 	int64   `gorm:"default:1;column: quantity"`
	Subtotal 	float64 `gorm:"default:0;column: subtotal"`
	CreatedAt 	time.Time
}
