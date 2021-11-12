package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./pic")
	err := router.Run(":8080")
	if err != nil {
		return 
	}
}