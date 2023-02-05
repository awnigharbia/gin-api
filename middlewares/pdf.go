package middlewares

import (
	"ginApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PDFOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is missing"})
			c.Abort()
			return
		}

		// Check file using checkFile function
		fileTypes := []string{".pdf"}
		fileSizeMin := int64(0)
		fileSizeMax := int64(50000000)
		result := utils.CheckFile(header.Filename, fileTypes, fileSizeMin, fileSizeMax)
		if result != header.Filename + " is a valid file." {
			c.JSON(http.StatusBadRequest, gin.H{"error": "PDF only, max size 50MB"})
			c.Abort()
			return
		}

		c.Next()
	}
}