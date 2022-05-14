package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc/metadata"

	common "github.com/Mr-Herod/CloudGamingDemo/Common"
	gaming "github.com/Mr-Herod/CloudGamingDemo/Gaming/gaming"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StartGameReq struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Gamename  string `json:"gamename"`
	ClientDes string `json:"clientDes"`
}

func HandleStartGame(c *gin.Context) {
	var reqBody StartGameReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	defer cancel()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("username", reqBody.Username, "nickname", reqBody.Nickname, "gamename", reqBody.Gamename),
	)

	rspStr, err := gamingHandler(ctx, reqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, rspStr)
}

func gamingHandler(ctx context.Context, req StartGameReq) (string, error) {
	clientDes := req.ClientDes
	log.Printf("HandleStartGame ctx:%+v", ctx)
	log.Printf("ClientDes: %v", clientDes[len(clientDes)-20:])
	// Set up a connection to the server.
	ip, port, _ := common.FindServer("Gaming")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		return "", fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	cli := gaming.NewGamingServiceClient(conn)
	rsp, err := cli.StartGame(ctx, &gaming.StartGameReq{
		Username:  req.Username,
		Nickname:  req.Nickname,
		Gamename:  req.Gamename,
		ClientDes: clientDes,
	})
	if err != nil || rsp.ErrCode != 0 {
		log.Printf("StartGame error: %v", err)
		return "", fmt.Errorf("StartGame error: %v", err)
	}
	log.Printf("ServerDes: %v", rsp.ServerDes[len(rsp.ServerDes)-20:])
	return rsp.ServerDes, nil
}
