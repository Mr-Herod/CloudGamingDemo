package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	account "github.com/Mr-Herod/CloudGamingDemo/Account/account"
	common "github.com/Mr-Herod/CloudGamingDemo/Common"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func HandleSignUp(c *gin.Context) {
	// Set up a connection to the server.
	ip, port, _ := common.FindServer("Account")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer conn.Close()
	cli := account.NewAccountServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	rsp, err := cli.UserRegister(ctx, &account.UserRegisterReq{
		Username: username,
		Password: password,
		Nickname: nickname,
	})
	if err != nil {
		log.Printf("UserRegister error: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	log.Printf("rsp: %v", rsp)
	c.SetCookie("username", username, 60*60*24, "/", "mrherod.cn", false, false)
	c.SetCookie("nickname", nickname, 60*60*24, "/", "mrherod.cn", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/center")
}

func HandleSignIn(c *gin.Context) {
	// Set up a connection to the server.
	ip, port, _ := common.FindServer("Account")
	conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer conn.Close()
	cli := account.NewAccountServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	username := c.PostForm("username")
	password := c.PostForm("password")
	rsp, err := cli.UserLogIn(ctx, &account.UserLogInReq{
		Username: username,
		Password: password,
	})
	if err != nil || rsp.ErrCode != 0 {
		log.Printf("UserLogin error: %v", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	log.Printf("rsp: %v", rsp)
	c.SetCookie("username", username, 60*60*24, "/", "mrherod.cn", false, false)
	c.SetCookie("nickname", rsp.Nickname, 60*60*24, "/", "mrherod.cn", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/center")
}
