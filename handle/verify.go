package handle

import (
	"easyPicServer/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var VerifiedMapMutex sync.Mutex
var VerifiedMap = make(map[string]int64)

type Verified struct {
	TimeStamp int64
	sCode     string
}

var TimeQueue = make([]Verified, 10)

func CleanTimeUpVerifiedInfo() {
	go func() {
		for {
			time.Sleep(time.Second * 10)
			cfg := config.GetGlobalConfig()
			now := time.Now().Unix()
			VerifiedMapMutex.Lock()
			for i, v := range TimeQueue {
				if v.TimeStamp+cfg.EffectiveDur > now {
					TimeQueue = TimeQueue[i:]
					break
				}
				delete(VerifiedMap, v.sCode)
			}
			VerifiedMapMutex.Unlock()
		}
	}()
}

func GenerateQuestion(context *gin.Context) {
	cfg := config.GetGlobalConfig()
	r := (rand.Int() % len(cfg.QaMap)) + 1
	question := cfg.QaMap[r]
	plaintext := fmt.Sprintf("%v|%v|%v", r, time.Now().Unix(), rand.Int63())
	ciphertext, err := cfg.Encryption.EncryptAndBase64([]byte(plaintext))
	if err != nil {
		context.String(http.StatusInternalServerError, "加密出现BUG，耶稣来了也进不去，我说的！")
		return
	}
	context.HTML(http.StatusOK, "security.html", gin.H{
		"title":    "Security",
		"label":    "You need answer the question!",
		"question": question.Question,
		"pic":      context.Param("pic"),
		"sCode":    ciphertext,
	})
}

func tableCheckMatch(sCode, answer, oldSCode string) bool {
	cfg := config.GetGlobalConfig()
	now := time.Now().Unix()
	VerifiedMapMutex.Lock()
	defer VerifiedMapMutex.Unlock()
	if oldTime, ok := VerifiedMap[oldSCode]; ok {
		return oldTime+cfg.EffectiveDur > now
	}
	plaintext, err := cfg.Encryption.DecryptFromBase64(sCode)
	if err != nil {
		return false
	}
	texts := strings.Split(string(plaintext), "|")
	timestamp, err := strconv.ParseInt(texts[2], 10, 0)
	if err != nil {
		return false
	}
	if timestamp+cfg.EffectiveDur < now {
		return false
	}
	r, err := strconv.Atoi(texts[0])
	if err != nil {
		return false
	}
	qa := cfg.QaMap[r]
	for _, a := range qa.Answers {
		if a == answer {
			VerifiedMap[sCode] = now
			TimeQueue = append(TimeQueue, Verified{TimeStamp: now, sCode: sCode})
			return true
		}
	}
	return false
}

func CheckAndGetPic(context *gin.Context) {
	sCode := context.Param("sCode")
	answer := context.PostForm("answer")
	oldSCode := context.PostForm("oldsCode")
	if tableCheckMatch(sCode, answer, oldSCode) {
		context.File("./pic/" + context.Param("pic"))
	} else {
		context.String(http.StatusNotFound, "你弄撒咧？")
	}
}
