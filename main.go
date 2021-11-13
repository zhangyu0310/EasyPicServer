package main

import (
	"easyPicServer/config"
	"easyPicServer/dumpserver"
	"flag"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

var (
	webPort  = flag.Int("web-port", 8888, "Port of web server")
	dumpPort = flag.Int("dump-port", 9999, "Port of dump server")
)

// cmdConfigSetToGlobal store command config to global config.
func cmdConfigSetToGlobal(cfg *config.Config) {
	cfg.WebPort = *webPort
	cfg.DumpPort = *dumpPort
}

func main() {
	help := flag.Bool("help", false, "show the usage")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	config.InitializeConfig(cmdConfigSetToGlobal)
	// Start dump server
	dumpserver.Run()
	// Start web server
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./pic")
	err := router.Run(":" + strconv.Itoa(config.GetGlobalConfig().WebPort))
	if err != nil {
		return
	}
}
