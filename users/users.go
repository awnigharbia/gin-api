package users

import (
	"fmt"
	"ginApp/database"
	"ginApp/middlewares"
	"ginApp/models"
	"ginApp/modules/auth"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Users(router *gin.Engine) {
	userGroup := router.Group("/users")
	authGroup := router.Group("/auth")
	userGroup.Use(middlewares.AuthMiddleware())

	userGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "users"})
	})

	userGroup.GET("/info", func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("AllYourBase"), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims);
		
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": err.Error()})
			return
		}

		convertedstr := fmt.Sprintf("%v", claims["id"])

		user, err := models.GetUserByID(database.Db, convertedstr)
		
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Fetched successfully", "user": user})
	})

	authGroup.POST("/register", auth.RegisterHandler)
}