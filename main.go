package main

import (
	"log"
	"os"

	localstorage "github.com/Harmeet10000/go-file-upload-api/localStorage"
	s3storage "github.com/Harmeet10000/go-file-upload-api/s3Storage"
	"github.com/gin-gonic/gin"
)

func main() {
	// Ensure upload directory exists
	if err := os.MkdirAll(localstorage.UploadDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create upload dir: %v", err)
	}
	router := gin.Default()

	// NOOB Mistake 1: Not setting the max upload size
	// Max upload size: 10MB
	router.MaxMultipartMemory = 10 << 20 // 10 MiB

	router.POST("/upload", localstorage.HandleUpload)
	router.GET("/uploads/:filename", localstorage.HandleGetFile)
	router.POST("/s3/upload", s3storage.HandleUpload)

	log.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}