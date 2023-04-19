package main

import (
	"github.com/iqbaludinm/mygram-api/app"
	"github.com/iqbaludinm/mygram-api/config"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// init gorm
	err = config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApp()
}
