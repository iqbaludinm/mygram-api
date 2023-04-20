package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/mygram-api/handler"
	"github.com/iqbaludinm/mygram-api/middlewares"
	
	_ "github.com/iqbaludinm/mygram-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterAPI(router *gin.Engine, server handler.HttpServer) *gin.Engine {
	
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", server.Login)
			auth.POST("/register", server.Register)
		}
		
		v1.Use(middlewares.Authentication())
		photo := v1.Group("/photos")
		{
			photo.POST("", server.CreatePhoto)
			photo.GET("", server.GetPhotos)
			photo.GET("/:photoId", server.GetPhotoById)
			photo.PUT("/:photoId", server.UpdatePhoto)
			photo.DELETE("/:photoId", server.DeletePhoto)
		}

		comment := v1.Group("/comments")
		{
			comment.POST("", server.CreateComment)
			comment.GET("", server.GetComments)
			comment.GET("/:commentId", server.GetCommentById)
			comment.PUT("/:commentId", server.UpdateComment)
			comment.DELETE("/:commentId", server.DeleteComment)
		}

		socmed := v1.Group("/social-medias")
		{
			socmed.POST("", server.CreateSocialMedia)
			socmed.GET("", server.GetSocialMedias)
			socmed.GET("/:socmedId", server.GetSocialMediaById)
			socmed.PUT("/:socmedId", server.UpdateSocialMedia)
			socmed.DELETE("/:socmedId", server.DeleteSocialMedia)
		}
	}
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
