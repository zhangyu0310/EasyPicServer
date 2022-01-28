package main

import (
	"context"
	"easyPicServer/config"
	"easyPicServer/encrypt/aes"
	"easyPicServer/handle"
	"easyPicServer/store/storagefactory"
	"easyPicServer/transmit"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	webPort      = flag.Int("web-port", 8888, "Port of web server")
	dumpPort     = flag.Int("dump-port", 9999, "Port of dump server")
	effectiveDur = flag.Int64("effective-duration", 300, "sCode effective duration (second)")
	qaTableFile  = flag.String("qa-table-file", "./QATable.xlsx", "QA Table file")
	keyFile      = flag.String("key-file", "./key", "Key file for encryption")
	dbType       = flag.String("db-type", "leveldb", "Type of database (Default: LevelDB)")
	dbAddr       = flag.String("db-addr", "./db", "Address of database")
	dbPassword   = flag.String("db-password", "", "Password of database")
	cleanUpCount = flag.Uint("clean-up-count", 15, "Count of post failed to clean up webhook.")
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
	cfg.DBType = *dbType
	cfg.DBAddr = *dbAddr
	cfg.DBPassword = *dbPassword
	cfg.CleanUpCount = uint32(*cleanUpCount)
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
	db, err := storagefactory.CreateStorage(cfg)
	if err != nil {
		log.Println("Create Storage failed.", err)
		os.Exit(1)
	}
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
	srv := &http.Server{Addr: ":" + strconv.Itoa(cfg.WebPort), Handler: router}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")
	if err := db.Close(); err != nil {
		log.Fatal("LevelDB close failed.", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
