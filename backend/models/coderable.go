package models

import "time"

type Coderable struct {
	CodeID   	uint
	Code     	Code
	OrderID		uint
	Order     	Order
	CreatedAt 	time.Time
}
