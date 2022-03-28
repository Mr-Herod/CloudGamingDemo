package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	common "github.com/Mr-Herod/CloudGamingDemo/Common"
	gaming "github.com/Mr-Herod/CloudGamingDemo/Gaming/gaming"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StartGameReq struct {
	Username  string `json:"username"`
	ClientDes string `json:"clientDes"`
}

func HandleStartGame(c *gin.Context) {
	var reqBody StartGameReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rspStr, err := gamingHandler(reqBody.Username, reqBody.ClientDes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, rspStr)
}

func gamingHandler(username, clientDes string) (string, error) {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rsp, err := cli.StartGame(ctx, &gaming.StartGameReq{
		Username:  username,
		ClientDes: clientDes,
	})
	if err != nil || rsp.ErrCode != 0 {
		log.Printf("StartGame error: %v", err)
		return "", fmt.Errorf("StartGame error: %v", err)
	}
	log.Printf("ServerDes: %v", rsp.ServerDes[len(rsp.ServerDes)-20:])
	return rsp.ServerDes, nil
}
