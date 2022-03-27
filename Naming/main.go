package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Mr-Herod/CloudGamingDemo/Naming/naming"

	"google.golang.org/grpc"
)

type serverInfo struct {
	Name string
	Ip   string
	Port int32
}

var (
	ServerMap = make(map[string]*serverInfo)
	Port      = flag.Int("Port", 10086, "The server Port")
)

type server struct {
	pb.UnimplementedNamingServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNamingServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) RegisterService(ctx context.Context, req *pb.RegisterServiceReq) (*pb.RegisterServiceRsp, error) {
	ServerMap[req.ServiceName] = &serverInfo{
		Name: req.ServiceName,
		Ip:   req.IP,
		Port: req.Port,
	}
	log.Printf("RegisterService: %+v succeeded!", ServerMap[req.ServiceName])
	return &pb.RegisterServiceRsp{
		ErrCode: 0,
		Msg:     "RegisterService succeeded!",
	}, nil
}

func (s *server) FindService(ctx context.Context, req *pb.FindServiceReq) (rsp *pb.FindServiceRsp, err error) {
	rsp = &pb.FindServiceRsp{}
	log.Printf("ServerMap: %+v !", ServerMap)
	log.Printf("Finding server: %v %+v!", req.ServiceName, ServerMap[req.ServiceName])
	info := ServerMap[req.ServiceName]
	rsp.ServiceName = info.Name
	rsp.IP = info.Ip
	rsp.Port = info.Port
	return
}
