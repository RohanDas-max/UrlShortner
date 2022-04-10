package database

import (
	"github.com/rohandas-max/urlShortner/pkg/models"
	"github.com/rohandas-max/urlShortner/pkg/platform"
)

func InsertUrl(r platform.UrlReqBody, randValue string) {
	dB.Create(&models.Url{
		LongUrl:  r.Url,
		ShortUrl: randValue,
	})
}

func GetUrl(url string) (v string) {
	dB.Table("urls").Select("long_url").Where("short_url = ?;", url).Scan(&v)
	return
}

func IsExist(url string) bool {
	var id int
	dB.Table("urls").Select("id").Where("long_url=?", url).Scan(&id)
	if id > 0 {
		return true
	}
	return false
}

func DropTable() {
	dB.Migrator().DropTable("urls")
}
