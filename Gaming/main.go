package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	common "github.com/Mr-Herod/CloudGamingDemo/Common"
	config "github.com/Mr-Herod/CloudGamingDemo/Gaming/config"
	pb "github.com/Mr-Herod/CloudGamingDemo/Gaming/gaming"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGamingServiceServer
}

func main() {
	log.SetFlags(log.Lshortfile | log.Llongfile) // set flags
	config.Init()
	InitGame()
	cfg := config.ServiceConf
	port := flag.Int("port", cfg.ListenPort, "The server port")
	flag.Parse()
	err := common.RegisterServer(cfg.ServiceName, "localhost", int32(*port))
	if err != nil {
		panic(fmt.Sprintf("RegisterServer falied,err:%v", err))
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGamingServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) StartGame(ctx context.Context, req *pb.StartGameReq) (rsp *pb.StartGameRsp, err error) {
	desChan := make(chan string)
	go RTC(ctx, req.ClientDes, desChan)
	serverDes := <-desChan

	return &pb.StartGameRsp{ServerDes: serverDes}, nil
}
