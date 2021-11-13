package dumpserver

import (
	"context"
	"easyPicServer/config"
	"easyPicServer/picdump"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

type dumpServer struct{}

func (ds *dumpServer) SendPic(ctx context.Context, req *picdump.PicRequest) (*picdump.PicReply, error) {
	var reply picdump.PicReply
	reply.Message = "Success."
	err := ioutil.WriteFile("./pic/"+req.GetPicName(), req.Pic, 0666)
	if err != nil {
		reply.Message = "Write pic file failed."
		return &reply, err
	}
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
	picdump.RegisterCourierServer(server, &dumpServer{})

	reflection.Register(server)

	go func() {
		err = server.Serve(listen)
		if err != nil {
			fmt.Println("Dump Server run failed.", err)
		}
	}()
}
