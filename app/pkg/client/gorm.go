package client

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewGormClient(connectionString string) (gormClient *gorm.DB, err error) {
	gormClient, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect  %v\n", err)
	}
	return gormClient, nil
}
