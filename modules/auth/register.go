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


func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func RegisterHandler(c *gin.Context) {
	var userData struct {
		Name string `form:"name" binding:"required"`
		Email string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
		FcmToken string `form:"fcm_token" binding:"required"`
	}


	if err := c.ShouldBind(&userData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	existingUser := models.User{Email: userData.Email}
	has, err := database.Db.Table("users").Get(&existingUser)
	
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	if has {
		c.JSON(400, gin.H{"error": "Email already exists"})
		return
	}


	hashedPassword, err := hashPassword(userData.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Hash error"})
		return
	}

	user := models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: hashedPassword,
		FcmToken: userData.FcmToken,
	}

	userID, err := database.Db.Table("users").Insert(&user)
	 
	 if err != nil {
		c.JSON(500, gin.H{"error": "Insertion error"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No id returned"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("AllYourBase"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}