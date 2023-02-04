package auth

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.POST("/login", LoginHandler)
	router.POST("/register", RegisterHandler)
}