package main

import (
	"easyPicServer/config"
	"easyPicServer/encrypt/aes"
	"easyPicServer/handle"
	"easyPicServer/store"
	ldb "easyPicServer/store/leveldb"
	"easyPicServer/transmit"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	webPort      = flag.Int("web-port", 8888, "Port of web server")
	dumpPort     = flag.Int("dump-port", 9999, "Port of dump server")
	effectiveDur = flag.Int64("effective-duration", 300, "sCode effective duration. (second)")
	qaTableFile  = flag.String("qa-table-file", "./QATable.xlsx", "QA Table file")
	keyFile      = flag.String("key-file", "./key", "Key file for encryption")
	dbPath       = flag.String("db-path", "./db", "DB file path")
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
	cfg.DBPath = *dbPath
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
	transmit.Run()
	// Init storage
	cfg := config.GetGlobalConfig()
	db, err := leveldb.OpenFile(cfg.DBPath, nil)
	if err != nil {
		fmt.Println("Initialize storage failed!", err)
		os.Exit(1)
	}
	store.InitializeStorage(&ldb.LevelDB{DB: db})
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
	router.GET("/register", handle.RegisterWeChatWebhook)
	router.POST("/result", handle.RegisterResult)
	handle.CleanTimeUpVerifiedInfo()
	err = router.Run(":" + strconv.Itoa(cfg.WebPort))
	if err != nil {
		return
	}
}
