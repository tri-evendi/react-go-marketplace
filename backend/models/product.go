package models

type Product struct {
	ID          uint   `gorm:"primary_key;NOT NULL AUTO_INCREMENT"`
	Name        string `gorm:"type:varchar(255);column: name;NOT NULL"`
	Description string `gorm:"column: description"`
	CategoryID  uint
	Category    Category
	Price       float64 `gorm:"default:0;column: price"`
	ImagePath   string  `gorm:"default:misc/bgimage.jpg"`
}
