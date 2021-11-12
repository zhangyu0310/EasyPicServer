package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	port = flag.Int("port", 8888, "Port of server.")
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./pic")
	err := router.Run(":" + strconv.Itoa(*port))
	if err != nil {
		return
	}
}
