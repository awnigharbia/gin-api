package auth

import (
	"ginApp/database"
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
	var user struct {
		Name string `form:"name" binding:"required"`
		Email string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
		FcmToken string `form:"fcm_token" binding:"required"`
	}


	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(user.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Hash error"})
		return
	}

	res, err := database.Db.Exec("INSERT INTO users (name, email, password, fcm_token) VALUES (?, ?, ?, ?)", user.Name, user.Email, hashedPassword, user.FcmToken); 
	 
	 if err != nil {
		c.JSON(500, gin.H{"error": "Insertion error"})
		return
	}

	userID, err := res.LastInsertId()

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