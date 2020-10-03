package main

import (
	"base-site-api/internal/app/http"
	"base-site-api/internal/log"
)

func main() {
	c, err := http.New()

	if err != nil {
		log.Fatal(err)
	}

	app := http.NewApp(c)
	log.Debug(c.Constants.ADDRESS)
	err = app.Listen(c.Constants.ADDRESS)

	if err != nil {
		log.Fatal(err)
	}
}
