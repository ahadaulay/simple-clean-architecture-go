package helpers

import (
	"belajar-clean/drivers"

	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	configDatabase := drivers.ConfigDB{
		DB_USERNAME: GetConfig("DB_USERNAME"),	
		DB_PASSWORD: GetConfig("DB_PASSWORD"),	
		DB_NAME: GetConfig("DB_NAME"),	
	}

	db := configDatabase.InitDB()

	err := drivers.DBMigration(db)

	if err != nil {
		panic(err)
	}

	return db
}