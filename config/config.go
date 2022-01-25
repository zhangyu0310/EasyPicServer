package config

import (
	"easyPicServer/encrypt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"sync/atomic"
)

// Config Configurations of server.
type Config struct {
	WebPort      int
	DumpPort     int
	EffectiveDur int64
	QaFilePath   string
	QaMap        map[int]QATable
	Encryption   encrypt.Encryption
	DBPath       string
}

type QATable struct {
	Id       int
	Question string
	Answers  []string
}

var (
	globalConf atomic.Value
)

// InitializeConfig initialize the global config handler.
func InitializeConfig(enforceCmdArgs func(*Config)) {
	cfg := Config{}
	// Use command config cover config file.
	enforceCmdArgs(&cfg)
	InitQaMapFromExcel(&cfg)
	StoreGlobalConfig(&cfg)
}

// GetGlobalConfig returns the global configuration for this server.
// It should store configuration from command line and configuration file.
// Other parts of the system can read the global configuration use this function.
func GetGlobalConfig() *Config {
	return globalConf.Load().(*Config)
}

// StoreGlobalConfig stores a new config to the globalConf. It mostly uses in the test to avoid some data races.
func StoreGlobalConfig(config *Config) {
	globalConf.Store(config)
}

func InitQaMapFromExcel(config *Config) {
	xlFile, err := xlsx.OpenFile(config.QaFilePath)
	if err != nil {
		log.Println("Open QA Table failed.", err)
		os.Exit(1)
	}
	sheet := xlFile.Sheets[0]
	if sheet == nil {
		log.Println("First sheet is nil.")
		os.Exit(1)
	}
	config.QaMap = make(map[int]QATable)
	for i, row := range sheet.Rows {
		if i == 0 {
			continue
		}
		id, err := row.Cells[0].Int()
		if err != nil {
			log.Println("Get Question ID failed.")
			os.Exit(1)
		}
		var answers []string
		for i := 2; i < len(row.Cells); i++ {
			answer := row.Cells[i].String()
			if answer != "" {
				answers = append(answers, answer)
			}
		}
		table := QATable{Id: id, Question: row.Cells[1].String(), Answers: answers}
		config.QaMap[id] = table
	}
}
