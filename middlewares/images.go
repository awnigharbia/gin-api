package middlewares

import (
	"ginApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImagesOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is missing"})
			c.Abort()
			return
		}

		// Check file using checkFile function
		fileTypes := []string{".jpg", ".jpeg", ".png", ".gif"}
		fileSizeMin := int64(0)
		fileSizeMax := int64(3000000)
		result := utils.CheckFile(header.Filename, fileTypes, fileSizeMin, fileSizeMax)
		if result != header.Filename + " is a valid file." {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Images only, max size 3MB"})
			c.Abort()
			return
		}

		c.Next()
	}
}