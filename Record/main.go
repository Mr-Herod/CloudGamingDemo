package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	common "github.com/Mr-Herod/CloudGamingDemo/Common"
	config "github.com/Mr-Herod/CloudGamingDemo/Record/config"
	pb "github.com/Mr-Herod/CloudGamingDemo/Record/record"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRecordServiceServer
}

var DB *sql.DB

const DefaultID = 0

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

func main() {
	log.SetFlags(log.Lshortfile | log.Llongfile) // set flags
	config.Init()
	initDB()
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
	pb.RegisterRecordServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SaveRecord(ctx context.Context, req *pb.SaveRecordReq) (rsp *pb.SaveRecordRsp, err error) {
	rsp = &pb.SaveRecordRsp{}
	_, err = DB.Exec(
		"INSERT INTO records (recordID,userID,userName,gameID,gameName,score,recordTime) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?)",
		DefaultID, DefaultID, req.Username, DefaultID, req.Gamename, req.Score, time.Now().Unix(),
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

func (s *server) GetRank(ctx context.Context, req *pb.GetRankReq) (rsp *pb.GetRankRsp, err error) {
	rsp = &pb.GetRankRsp{}
	defer func() {
		if err != nil {
			rsp.ErrCode = -1
			rsp.Msg = fmt.Sprintf("GetRank error,err: %v", err)
		}
	}()
	var rows *sql.Rows
	if req.Username != "" && req.Gamename != "" {
		rows, err = DB.Query("SELECT userName,gameName,score,recordTime "+
			"FROM records WHERE userName=? AND gameName=? ORDER BY recordTime DESC",
			req.Username, req.Gamename)
	} else if req.Username != "" {
		rows, err = DB.Query("SELECT userName,gameName,score,recordTime "+
			"FROM records WHERE userName=? ORDER BY recordTime DESC", req.Username)
	} else if req.Gamename != "" {
		rows, err = DB.Query("SELECT userName,gameName,score,recordTime "+
			"FROM records WHERE gameName=? ORDER BY recordTime DESC", req.Gamename)
	} else {
		rows, err = DB.Query("SELECT userName,gameName,score,recordTime " +
			"FROM records ORDER BY recordTime DESC")
	}
	if err != nil {
		log.Fatal(err)
		return
	}
	records := []*pb.PlayRecord{}
	for rows.Next() {
		record := pb.PlayRecord{}
		if err = rows.Scan(&record.Username, &record.Gamename, &record.Score, &record.Time); err != nil {
			log.Fatal(err)
			return
		}
		records = append(records, &record)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return
	}
	rsp.Records = records
	return rsp, nil
}
