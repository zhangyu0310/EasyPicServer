package handle

import (
	"easyPicServer/store"
	"encoding/binary"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RegisterWeChatWebhook(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", gin.H{})
}

func RegisterResult(context *gin.Context) {
	webhook := context.PostForm("webhook")
	if !strings.HasPrefix(webhook, "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=") {
		context.String(http.StatusOK, "企业微信webhook有误，查证后填写！"+
			"(应当为'https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key='开头)")
		return
	}
	s := store.GetStorage()
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, 0)
	err := (*s).Set([]byte(webhook), b, nil)
	if err != nil {
		context.String(http.StatusOK, "注册失败... %v", err)
		return
	}
	context.String(http.StatusOK, "注册成功！")
}
