package database

import "github.com/rohandas-max/urlShortner/pkg/models"

func InsertUrl() {
	dB.Create(&models.Url{})
}
