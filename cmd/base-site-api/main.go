package main

import (
	"base-site-api/internal/app"
	"base-site-api/internal/app/config"
	"base-site-api/internal/database"
	"base-site-api/internal/log"
	"os"
)

func main() {
	// Load configurations
	c, err := config.Load()
	if err != nil {
		// Error when loading the configurations
		panic("An error occurred while loading the configurations: " + err.Error())
	}

	a := app.New(&c)
	err = a.Listen(c.App.Listen)

	if err != nil {
		appErr := a.Shutdown()
		if appErr != nil {
			log.Debugf("Problem with closing app %s ", err)
		}

		dbErr := database.Close()

		if dbErr != nil {
			log.Debugf("Problem with closing database %s", err)
		}

		// Return with corresponding exit code
		if dbErr != nil || appErr != nil {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
}
