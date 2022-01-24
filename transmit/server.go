package transmit

import (
	"bytes"
	"context"
	"easyPicServer/config"
	"easyPicServer/store"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
)

type dumpServer struct{}

func (ds *dumpServer) SendPic(_ context.Context, req *PicRequest) (*PicReply, error) {
	var reply PicReply
	reply.Message = "Success."
	err := ioutil.WriteFile("./pic/"+req.GetPicName(), req.Pic, 0666)
	if err != nil {
		reply.Message = "Write pic file failed."
		return &reply, err
	}
	return &reply, nil
}

func (ds *dumpServer) SendSuTu(_ context.Context, req *SeTuRequest) (*SeTuReply, error) {
	txtMsg := BotMsgReq{MsgType: BotMsgText,
		Text: &Text{Content: req.Url}}
	article := Article{
		Title:       req.Title,
		Description: req.Desc,
		Url:         req.OriginalUrl,
		Picurl:      req.OriginalUrl}
	var articles []Article
	articles = append(articles, article)
	newsMsg := BotMsgReq{MsgType: BotMsgNews,
		News: &News{
			Articles: articles}}
	picMsg := BotMsgReq{MsgType: BotMsgImage,
		Image: &Image{Base64: req.PicBase64, Md5: req.PicMd5}}
	for it := (*store.GetStorage()).Iterator(); it.Valid(); it.Next() {
		err := postSetuToWeChat(string(it.Key()), txtMsg)
		if err != nil {
			fmt.Println("Post txt msg failed.")
		}
		err = postSetuToWeChat(string(it.Key()), newsMsg)
		if err != nil {
			fmt.Println("Post txt msg failed.")
		}
		err = postSetuToWeChat(string(it.Key()), picMsg)
		if err != nil {
			fmt.Println("Post txt msg failed.")
		}
	}
	var reply SeTuReply
	reply.ErrMessage = "Success."
	return &reply, nil
}

func Run() {
	cfg := config.GetGlobalConfig()
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.DumpPort))
	if err != nil {
		fmt.Println("Dump server start failed.", err)
		os.Exit(1)
	}
	options := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(31457280),
		grpc.MaxSendMsgSize(31457280),
	}
	server := grpc.NewServer(options...)
	ds := dumpServer{}
	RegisterPicCourierServer(server, &ds)
	RegisterSetuCourierServer(server, &ds)

	reflection.Register(server)

	go func() {
		err = server.Serve(listen)
		if err != nil {
			fmt.Println("Dump Server run failed.", err)
		}
	}()
}

// postSetuToWeChat post setu to WeChat
func postSetuToWeChat(wechatUrl string, post BotMsgReq) (err error) {
	postStr, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Json marshal post failed.", err)
		return
	}
	respPost, err := http.Post(wechatUrl, "application/json", bytes.NewBuffer(postStr))
	if err != nil {
		fmt.Println("Post to wechat failed", err)
		return
	}
	msg, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(msg))
	_ = respPost.Body.Close()
	return
}
