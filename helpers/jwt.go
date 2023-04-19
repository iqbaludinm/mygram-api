package helpers

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		log.Fatalln(err)
	}
	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error){
	errResponse := errors.New("Failed to Verify Token")
	bearerToken := c.Request.Header.Get("Authorization")
	isBearer := strings.HasPrefix(bearerToken, "Bearer")
	if !isBearer {
		return nil, errResponse
	}
	strToken := strings.Split(bearerToken, " ")[1]
	token, _ := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // check signature method
		if !ok {
			return nil, errResponse
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	
	// validate mapClaims
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok && token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}