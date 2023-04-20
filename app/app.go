package app

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/mygram-api/config"
	"github.com/iqbaludinm/mygram-api/handler"
	"github.com/iqbaludinm/mygram-api/repositories"
	routesV1 "github.com/iqbaludinm/mygram-api/routes"
	"github.com/iqbaludinm/mygram-api/services"
)

var router = gin.New()

func StartApp() {
	repo := repositories.NewRepo(config.GORM.DB)
	service := services.NewService(repo)
	server := handler.NewHttpServer(service)

	router.Use(cors.Default())
	routesV1.RegisterAPI(router, server)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))

}