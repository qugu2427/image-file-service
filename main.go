package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func getMime(fileName string) string {
	mime := ""
	for i := len(fileName) - 1; i >= 0; i-- {
		if fileName[i] == '.' {
			break
		}
		mime = string(fileName[i]) + mime
	}
	return strings.ToLower(mime)
}

func isAllowedMime(mime string) bool {
	for i := 0; i < len(allowedMimes); i++ {
		if allowedMimes[i] == mime {
			return true
		}
	}
	return false
}

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if file.Size > maxSize {
		c.Status(413)
		return
	}
	if err != nil {
		c.Status(500)
		return
	}
	mime := getMime(file.Filename)
	if !isAllowedMime(mime) {
		c.Status(400)
		return
	}
	storageFileName := fmt.Sprintf("%v%v.%s", file.Size, time.Now().UnixNano(), getMime(file.Filename))
	c.SaveUploadedFile(file, uploadPath+storageFileName)
	c.String(201, storageFileName)
}

func deleteHandler(c *gin.Context) {
	fileName := c.PostForm("fileName")
	if strings.Contains(fileName, "/") || strings.Contains(fileName, "\\") {
		c.Status(400)
		return
	}
	os.Remove(uploadPath + fileName)
	c.Status(200)
}

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile(uploadPath, false)))

	router.POST("/upload", uploadHandler)
	router.DELETE("/delete", deleteHandler)

	router.Run(":3000")
}
