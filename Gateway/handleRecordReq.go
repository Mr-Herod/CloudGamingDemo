package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	common "github.com/Mr-Herod/CloudGamingDemo/Common"
	record "github.com/Mr-Herod/CloudGamingDemo/Record/record"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func HandleGetRank(c *gin.Context) {
	// Set up a connection to the server.
	ip, port, _ := common.FindServer("Record")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer conn.Close()
	cli := record.NewRecordServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	username := c.PostForm("username")
	gamename := c.PostForm("gamename")
	log.Printf("HandleGetRank username:%s\n", username)
	rsp, err := cli.GetRank(ctx, &record.GetRankReq{
		Username: username,
		Gamename: gamename,
	})
	if err != nil {
		log.Printf("GetRank error: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	log.Printf("rsp: %v", rsp)
	c.HTML(http.StatusOK, "rank.html", gin.H{"records": rsp.Records})
}

func HandleGetRank1(c *gin.Context) {
	// Set up a connection to the server.
	ip, port, _ := common.FindServer("Record")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer conn.Close()
	cli := record.NewRecordServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	username, err := c.Cookie("username")
	log.Printf("HandleGetRank username:%s\n", username)
	rsp, err := cli.GetRank(ctx, &record.GetRankReq{
		Username: username,
	})
	if err != nil {
		log.Printf("GetRank error: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	log.Printf("rsp: %v", rsp)
	c.HTML(http.StatusOK, "rank.html", gin.H{"records": rsp.Records})
}
