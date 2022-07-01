package config

import (
	"time"
)


func ParseStringToTime(stringtime string) time.Time {
	parsedDate, err := time.Parse("1/2/2006", stringtime)
	if err != nil {
		panic(err)
	}

	return parsedDate
}
