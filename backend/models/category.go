package models

type Category struct {
	ID          uint   `gorm:"primary_key;NOT NULL AUTO_INCREMENT"`
	Name        string `gorm:"type:varchar(255);column: name;NOT NULL"`
	Description string `gorm:"column: description"`
}
