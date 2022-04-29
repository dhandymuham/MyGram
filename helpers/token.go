package helpers

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "secret1234"

func GenerateToken(id uint, email string, username string) string {
	claims := jwt.MapClaims{
		"id":       id,
		"email":    email,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func ValidToken(c *gin.Context) (interface{}, error) {
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, fmt.Errorf("sign in")
	}

	Tokenstring := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(Tokenstring, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("sign in")

		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, fmt.Errorf("sign in")

	}
	return token.Claims.(jwt.MapClaims), nil
}
