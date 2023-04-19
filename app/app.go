package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/mygram-api/config"
	"github.com/iqbaludinm/mygram-api/handler"
	"github.com/iqbaludinm/mygram-api/repositories"
	"github.com/iqbaludinm/mygram-api/routes"
	"github.com/iqbaludinm/mygram-api/services"
)

var router = gin.New()

func StartApp() {
	repo := repositories.NewRepo(config.GORM.DB)
	service := services.NewService(repo)
	server := handler.NewHttpServer(service)

	routes.RegisterAPI(router, server)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))

}