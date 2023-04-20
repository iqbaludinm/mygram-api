package middlewares

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/mygram-api/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			helpers.Unauthorized(c, "Unauthenticated", err.Error())
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}