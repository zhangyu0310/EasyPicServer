package main

import (
	"easyPicServer/config"
	"easyPicServer/dumpserver"
	"easyPicServer/encrypt/aes"
	"easyPicServer/handle"
	"flag"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	webPort  = flag.Int("web-port", 8888, "Port of web server")
	dumpPort = flag.Int("dump-port", 9999, "Port of dump server")
	effectiveDur = flag.Int64("effective-duration", 300, "sCode effective duration. (second)")
	qaTableFile = flag.String("qa-table-file", "./QATable.xlsx", "QA Table file")
	keyFile = flag.String("key-file", "./key", "Key file for encryption")
)

// cmdConfigSetToGlobal store command config to global config.
func cmdConfigSetToGlobal(cfg *config.Config) {
	cfg.WebPort = *webPort
	cfg.DumpPort = *dumpPort
	cfg.EffectiveDur = *effectiveDur
	cfg.QaFilePath = *qaTableFile
	key, err := os.ReadFile(*keyFile)
	if err != nil {
		key = []byte("Golang is the best language!@#$%")
	}
	cfg.Encryption = &aes.Aes{PrivateKey: key}
}

func main() {
	help := flag.Bool("help", false, "show the usage")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	config.InitializeConfig(cmdConfigSetToGlobal)
	// Set random seed
	rand.Seed(time.Now().UnixNano())
	// Start dump server
	dumpserver.Run()
	// Start web server
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.LoadHTMLGlob("view/*")
	group := router.Group("/security")
	{
		group.GET("/:pic", handle.GenerateQuestion)
		group.POST("/:pic/:sCode", handle.CheckAndGetPic)
	}
	handle.CleanTimeUpVerifiedInfo()
	err := router.Run(":" + strconv.Itoa(config.GetGlobalConfig().WebPort))
	if err != nil {
		return
	}
}
