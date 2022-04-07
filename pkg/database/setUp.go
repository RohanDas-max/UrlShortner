package database

import (
	"github.com/rohandas-max/urlShortner/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dB *gorm.DB

func SetUp(dsn string) error {
	var err error
	dB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	dB.AutoMigrate(&models.Url{})
	return nil
}
