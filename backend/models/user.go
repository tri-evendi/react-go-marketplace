package models

type User struct {
	ID        uint   `gorm:"primary_key;NOT NULL AUTO_INCREMENT"`
	FirstName string `gorm:"type:varchar(255);column: firstName;NOT NULL"`
	LastName  string `gorm:"type:varchar(255);column: lastName;NOT NULL"`
	Email     string `gorm:"UNIQUE; NOT NULL"`
	Password  []byte
	Role      string `gorm:"type:varchar(100);column: role;NOT NULL"`
}
