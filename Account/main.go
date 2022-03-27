package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Mr-Herod/CloudGamingDemo/Account/account"
	config "github.com/Mr-Herod/CloudGamingDemo/Account/config"
	common "github.com/Mr-Herod/CloudGamingDemo/Common"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedAccountServiceServer
}

var DB *sql.DB

func main() {
	config.Init()
	initDB()
	port := flag.Int("port", config.ServiceConf.ListenPort, "The server port")
	flag.Parse()
	err := common.RegisterServer("Account", "localhost", int32(*port))
	if err != nil {
		panic(fmt.Sprintf("RegisterServer falied,err:%v", err))
	}
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

func initDB() {
	db, err := sql.Open("mysql", config.ServiceConf.DBSource)
	if err != nil {
		panic(err)
	}
	log.Println("Open Database succeeded!")
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	DB = db
}

func (s *server) UserRegister(ctx context.Context, req *pb.UserRegisterReq) (rsp *pb.UserRegisterRsp, err error) {
	rsp = &pb.UserRegisterRsp{}
	_, err = DB.Exec(
		"INSERT INTO users (userID,userName,password,nickName,bestRank) VALUES (?, ?, ?, ?, ?)",
		0, req.Username, req.Password, req.Nickname, -1,
	)
	if err != nil {
		log.Printf("DB.Exec error,err:%v", err)
		rsp.ErrCode = -1
		rsp.Msg = fmt.Sprintf("DB.Exec error,err:%v", err)
		return
	}
	log.Printf("UserRegister succseeded!")
	rsp.ErrCode = 0
	rsp.Msg = fmt.Sprintf("UserRegister succeeded!")
	return
}

func (s *server) UserLogIn(ctx context.Context, req *pb.UserLogInReq) (rsp *pb.UserLogInRsp, err error) {
	rsp = &pb.UserLogInRsp{}
	defer func() {
		if err != nil {
			rsp.ErrCode = -1
			rsp.Msg = fmt.Sprintf("UserLogIn error,err:%v", err)
		} else {
			rsp.ErrCode = 0
			ctx = context.WithValue(ctx, "username", req.Username)
			rsp.Msg = fmt.Sprintf("UserLogIn succeeded")
		}
	}()
	rows, err := DB.Query("SELECT password FROM users WHERE userName = ?", req.Username)
	if err != nil {
		log.Fatal(err)
		return
	}
	var password string
	for rows.Next() {
		if err = rows.Scan(&password); err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("%s’s password %d\n", req.Username, password)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return
	}
	if password != req.Password {
		err = fmt.Errorf("password not match!")
	}
	return
}

func (s *server) UserLogOut(ctx context.Context, req *pb.UserLogOutReq) (*pb.UserLogOutRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogOut not implemented")
}
