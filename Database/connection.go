package database

import (
	models "UBookTsk/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=aqkhan88 dbname=Bookdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the Book model
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
