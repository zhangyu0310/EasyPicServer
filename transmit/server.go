package transmit

import (
	"bytes"
	"context"
	"easyPicServer/config"
	"easyPicServer/store"
	"encoding/binary"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
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
	s := *store.GetStorage()
	it := s.Iterator()
	for ; it.Valid(); it.Next() {
		webhook := string(it.Key())
		value := binary.BigEndian.Uint32(it.Value())
		oldV := value
		err := postSetuToWeChat(webhook, txtMsg)
		if err != nil {
			value++
			log.Println("Post txt msg failed.", webhook)
		}
		err = postSetuToWeChat(webhook, newsMsg)
		if err != nil {
			value++
			log.Println("Post news msg failed.", webhook)
		}
		err = postSetuToWeChat(webhook, picMsg)
		if err != nil {
			value++
			log.Println("Post image msg failed.", webhook)
		}
		if oldV != value {
			v := make([]byte, 4)
			binary.BigEndian.PutUint32(v, value)
			if err := s.Set(it.Key(), v, nil); err != nil {
				log.Println("Set post failure value failed.")
			}
		}
	}
	it.Release()
	var reply SeTuReply
	reply.ErrMessage = "Success."
	return &reply, nil
}

func Run() {
	cfg := config.GetGlobalConfig()
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.DumpPort))
	if err != nil {
		log.Println("Dump server start failed.", err)
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
			log.Println("Dump Server run failed.", err)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Hour * 3)
			checkInvalidWebhook()
		}
	}()
}

// postSetuToWeChat post setu to WeChat
func postSetuToWeChat(wechatUrl string, post BotMsgReq) (err error) {
	postStr, err := json.Marshal(post)
	if err != nil {
		log.Println("Json marshal post failed.", err)
		return
	}
	respPost, err := http.Post(wechatUrl, "application/json", bytes.NewBuffer(postStr))
	if err != nil {
		log.Println("Post to wechat failed", err)
		return
	}
	msg, err := ioutil.ReadAll(respPost.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(msg))
	_ = respPost.Body.Close()
	return
}

func checkInvalidWebhook() {
	cfg := config.GetGlobalConfig()
	s := *store.GetStorage()
	it := s.Iterator()
	for ; it.Valid(); it.Next() {
		txtMsg := BotMsgReq{MsgType: BotMsgText,
			Text: &Text{Content: "Webhook发送失败率过高，已经删除该记录。"}}
		value := binary.BigEndian.Uint32(it.Value())
		if value > cfg.CleanUpCount {
			if err := s.Delete(it.Key()); err != nil {
				log.Println("Delete Invalid Webhook failed.", err)
			}
			if err := postSetuToWeChat(string(it.Value()), txtMsg); err != nil {
				log.Println("Post Delete Invalid Webhook msg failed.", err)
			}
		}
	}
	it.Release()
}