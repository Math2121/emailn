package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {

	connectionConfig := "host=localhost user=root password=root dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connectionConfig), &gorm.Config{})

	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
