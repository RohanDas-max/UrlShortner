package main

import (
	"log"

	"github.com/rohandas-max/urlShortner/pkg/database"
	"github.com/rohandas-max/urlShortner/pkg/router"
)

func main() {
	dsn := "host=localhost user=postgres password=abc@123 dbname=urlshortner port=4000 sslmode=disable"
	if err := database.SetUp(dsn); err != nil {
		log.Fatal(err)
	}
	router := router.Router()
	router.Run()
}
