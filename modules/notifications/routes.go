package notifications

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/notifications", NotificationsListener)
	router.GET("/notifications/user", NotificationsListener)
}