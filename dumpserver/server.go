package dumpserver

import (
	"context"
	"easyPicServer/config"
	"easyPicServer/picdump"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"strconv"
)

type dumpServer struct{}

func (ds *dumpServer) SendPic(ctx context.Context, req *picdump.PicRequest) (*picdump.PicReply, error) {
	var reply picdump.PicReply
	reply.Message = "Success."
	picFile, err := os.OpenFile("./pic/"+req.GetPicName(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		reply.Message = "Open pic file failed."
		return &reply, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(picFile)
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
