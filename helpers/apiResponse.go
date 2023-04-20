package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = "record not found"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OkWithData(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	obj := Response{
		Status:  http.StatusCreated,
		Message: message,
		Data:    data,
	}

	c.JSON(http.StatusCreated, obj)
}

func BadRequest(c *gin.Context, message string, data ...interface{}) {
	obj := Response{
		Status:  http.StatusBadRequest,
		Message: message,
	}

	if len(data) > 0 {
		obj.Data = data[0]
	}

	c.JSON(http.StatusBadRequest, obj)
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  http.StatusNotFound,
		Message: message,
		Data:    nil,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Status:  http.StatusInternalServerError,
		Message: message,
		Data:    "Internal Server Error",
	})
}

func Unauthorized(c *gin.Context, message string, err interface{}) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Status:  http.StatusUnauthorized,
		Message: message,
		Data:    err,
	})
	// c.JSON(http.StatusUnauthorized, Response{
	// 	Status:  http.StatusUnauthorized,
	// 	Message: message,
	// 	Data:    err,
	// })
}
