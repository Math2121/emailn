package database

import (
	"emailn/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func NewDb() *gorm.DB {

	connectionConfig := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connectionConfig), &gorm.Config{})

	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
