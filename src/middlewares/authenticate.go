package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tianmarillio/channela-backend/config"
	"github.com/tianmarillio/channela-backend/src/models"
)

func Authenticate(c *gin.Context) {
	var user *models.User

	// get userToken from headers
	reqToken := c.Request.Header.Get("Authorization")
	userToken := strings.Split(reqToken, "Bearer ")[1]

	// decode token
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// find user by id on token
		config.DB.Where("id = ?", claims["user_id"]).First(&user)
	} else {
		fmt.Println(err)
		c.AbortWithError(http.StatusUnauthorized, gin.Error{})
		return
	}

	// embed user to context
	c.Set("user", user)

	// continue
	c.Next()
}
