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
	serverMap = make(map[string]serverInfo)
	port      = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedNamingServiceServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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
	serverMap[req.ServiceName] = serverInfo{
		Name: req.ServiceName,
		Ip:   req.IP,
		Port: req.Port,
	}
	return nil, nil
}

func (s *server) FindService(ctx context.Context, req *pb.FindServiceReq) (rsp *pb.FindServiceRsp, err error) {
	info := serverMap[req.ServiceName]
	rsp.ServiceName = info.Name
	rsp.IP = info.Ip
	rsp.Port = info.Port
	return
}
