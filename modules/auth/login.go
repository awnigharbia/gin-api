package auth

import (
	"ginApp/database"
	"ginApp/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func LoginHandler(c *gin.Context) {
	var loginData struct {
		Email string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	var user models.User

	has, err := database.Db.Table("users").Where("email = ?", loginData.Email).Get(&user)

	if err != nil {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	if !has {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	if err := comparePassword(user.Password, loginData.Password); err != nil {
		c.JSON(401, gin.H{"error": "Wrong password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("AllYourBase"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}
