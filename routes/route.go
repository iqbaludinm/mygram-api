package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/mygram-api/handler"
	"github.com/iqbaludinm/mygram-api/middlewares"
)

func RegisterAPI(r *gin.Engine, server handler.HttpServer) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", server.Login)
		auth.POST("/register", server.Register)
	}

	// r.GET("/mygram-docs")

	v1 := r.Group("/api/v1")
	{
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

}
