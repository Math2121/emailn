package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {

	connectionConfig := "host=localhost user=root password=root dbname=postgres port=55000 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connectionConfig), &gorm.Config{})

	if err != nil{
		panic("fail to connect to database")
	}

	return db


}