package upload

import (
	"ginApp/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/upload",middlewares.PDFOnly(), UploadHandler)
}