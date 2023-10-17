package drivers

import (
	"belajar-clean/models/model"

	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) error {
	err := db.AutoMigrate(model.User{})

	if err != nil {
		return err
	}

	return nil
}