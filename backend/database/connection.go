package database

import (
	"backend/config"
	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB Connection global object
var DB *gorm.DB

func Connect() {
	// Added Params to parse time and set local time
	dsn := config.SqlUserName + ":" + config.SqlPassword + "@/" + config.SqlDatabaseName + "?parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("Can't connect database!")
	} else {
		println("Connection successful!")
		DB = connection
		var Found = connection.Migrator().HasTable(&models.User{})
		if !Found {
			RegenerateSchema()
		} else {
			println("Schema already exists!")
			// RegenerateSchema()
		}
	}
}

func RegenerateSchema() {
	// Added Params to parse time and set local time
	dsn := config.SqlUserName + ":" + config.SqlPassword + "@/" + config.SqlDatabaseName + "?parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn))
	// Drop all tables
	err = connection.Migrator().DropTable(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Code{},
		&models.Order{},
		&models.Orderable{},
		&models.Coderable{},
		&models.Payment{})

	if err != nil {
		panic("Can't drop tables!")
	} else {
		println("Tables dropped!")
		// Migrate the schema
		err = connection.AutoMigrate(
			&models.User{},
			&models.Category{},
			&models.Product{},
			&models.Code{},
			&models.Order{},
			&models.Orderable{},
			&models.Coderable{},
			&models.Payment{})

		if err != nil {
			println("Can't migrate schema!")
		} else {
			// To be run only when the database is empty
			var Found bool
			DB.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = 1) AS found").Scan(&Found) // check if row exists
			if !Found {
				println("Running data seeder...")
				PopulateDB()
				println("Data seeder successful!")
			}
		}
	}
}
