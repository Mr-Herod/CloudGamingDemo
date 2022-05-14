package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Mr-Herod/CloudGamingDemo/Naming/naming"

	"github.com/go-redis/redis/v8"
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
	RedisCli  = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
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
	getServer("")
}

func (s *server) RegisterService(ctx context.Context, req *pb.RegisterServiceReq) (*pb.RegisterServiceRsp, error) {
	ServerMap[req.ServiceName] = &serverInfo{
		Name: req.ServiceName,
		Ip:   req.IP,
		Port: req.Port,
	}
	setServer(req.ServiceName, *ServerMap[req.ServiceName])
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
	if info == nil {
		info = getServer(req.ServiceName)
		if info == nil {
			rsp.ErrCode = -1
			err = fmt.Errorf("can't find server:%v", req.ServiceName)
			return
		}
	}
	rsp.ServiceName = info.Name
	rsp.IP = info.Ip
	rsp.Port = info.Port
	return
}

func setServer(serverName string, info serverInfo) {
	ctx := context.Background()
	infoByte, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("json.Marshal(info) error,err:%v", err)
		return
	}
	err = RedisCli.HSet(ctx, "serverMap", serverName, string(infoByte)).Err()
	if err != nil {
		fmt.Printf("rdb.Set error,serverName:%v, info:%v,err:%v", serverName, info, err)
		return
	}
	log.Printf("Set service redis: %+v succeeded!", ServerMap[serverName])
}

func getServer(serverName string) (info *serverInfo) {
	info = &serverInfo{}
	ctx := context.Background()
	if serverName == "" {
		val, err := RedisCli.HGetAll(ctx, "serverMap").Result()
		if err == redis.Nil {
			fmt.Println("serverMap does not exist")
		} else if err != nil {
			fmt.Printf(" rdb.HGetAll error,err:%v\n", err)
			return nil
		}
		fmt.Println("serverMap in redis:", val)
	} else {
		val, err := RedisCli.HGet(ctx, "serverMap", serverName).Result()
		if err == redis.Nil {
			fmt.Printf("serverName:%v does not exist\n", serverName)
		} else if err != nil {
			fmt.Printf(" rdb.HGetAll error,err:%v\n", err)
			return nil
		}
		_ = json.Unmarshal([]byte(val), info)
		return info
	}
	return nil
}
