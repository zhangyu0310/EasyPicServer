package handle

import (
	"easyPicServer/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterWeChatWebhook(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", gin.H{})
}

func RegisterResult(context *gin.Context) {
	webhook := context.PostForm("webhook")
	s := store.GetStorage()
	err := (*s).Set([]byte(webhook), []byte(""), nil)
	if err != nil {
		context.String(http.StatusOK, "注册失败... %v", err)
		return
	}
	context.String(http.StatusOK, "注册成功！")
}
