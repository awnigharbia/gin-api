package upload

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	filename := header.Filename
	folder := "files"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0755)
	}

	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	ext := filepath.Ext(header.Filename)
	i := 1
	newFileName := filename + ext
	for {
		if _, err := os.Stat(filepath.Join(folder, newFileName)); os.IsNotExist(err) {
			break
		}
		newFileName = filename + "-" + strconv.Itoa(i) + ext
		i++
	}

	out, err := os.Create(filepath.Join(folder, newFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, "Uploaded")
}

