package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Mr-Herod/CloudGamingDemo/Account/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type server struct {
	pb.UnimplementedAccountServiceServer
}

func main() {
	naming
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) UserRegister(ctx context.Context, req *pb.UserRegisterReq) (*pb.UserRegisterRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}

func (s *server) UserLogIn(ctx context.Context, req *pb.UserLogInReq) (*pb.UserLogInRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogIn not implemented")
}

func (s *server) UserLogOut(ctx context.Context, req *pb.UserLogOutReq) (*pb.UserLogOutRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogOut not implemented")
}

func RegisterServer() {

}
