package unittest

import (
	"backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestConnection(t *testing.T) {
	dsn := config.SqlUserName + ":" + config.SqlPassword + "@/" + config.SqlDatabaseName + "?parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		t.Errorf("Connection to DB Failed")
		t.Fail()
	}

}
