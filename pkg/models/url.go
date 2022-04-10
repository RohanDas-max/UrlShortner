package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	LongUrl  string `json:"longurl"`
	ShortUrl string `json:"shorturl"`
}
