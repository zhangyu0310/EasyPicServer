package storagefactory

import (
	"easyPicServer/config"
	"easyPicServer/store"
	"easyPicServer/store/leveldb"
	"easyPicServer/store/redis"
	"errors"
	rds "github.com/go-redis/redis"
	"strings"
)

func CreateStorage(cfg *config.Config) (db store.Storage, err error) {
	switch strings.ToLower(cfg.DBType) {
	case "leveldb":
		db = &leveldb.LevelDB{}
		err = db.Init(cfg.DBAddr, nil)
	case "redis":
		db = &redis.Redis{}
		err = db.Init(&rds.Options{Addr: cfg.DBAddr, Password: cfg.DBPassword, DB: 0})
	default:
		return nil, errors.New("check the db name or implementation")
	}
	store.InitializeStorage(db)
	return
}
