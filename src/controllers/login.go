package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tianmarillio/channela-backend/config"
	"github.com/tianmarillio/channela-backend/src/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginController struct{}

var LoginController loginController

func (ctr loginController) Login(c *gin.Context) {
	// get & bind request body
	var body *AuthBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// get user by username
	var user models.User
	result := config.DB.
		Where("username = ?", body.Username).
		First(&user)
	if result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	// verify password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(body.Password),
	); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"expired_at": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	c.JSON(http.StatusAccepted, gin.H{
		"userToken": tokenString,
	})
}
