package main

import (
	"github.com/iqbaludinm/mygram-api/app"
	"github.com/iqbaludinm/mygram-api/config"
)

func init() {
	// init gorm
	err := config.InitGorm()
	if err != nil {
		panic(err)
	}
}

// @securitydefinitions.apikey	ApiKeyAuth
// @in header
// @name Authorization

// @title           MyGram API
// @version         1.0
// @description     This is a sample service for social-media.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  iqbaludinm14@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      mygram-api-production-ef65.up.railway.app
// @BasePath  /api/v1
func main() {
	app.StartApp()
}
