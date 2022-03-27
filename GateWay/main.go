package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Mr-Herod/CloudGamingDemo/Account/account"
	common "github.com/Mr-Herod/CloudGamingDemo/Common"
	"github.com/Mr-Herod/CloudGamingDemo/Gateway/config"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config.Init()
	resourcePath := config.ServiceConf.ResourcePath
	// 1.创建路由
	r := gin.Default()
	r.POST("/sendDes", HandleDes)
	r.POST("/signup", HandleSignUp)
	r.POST("/signin", HandleSignIn)
	r.Static("/js", resourcePath+"js")
	r.Static("/css", resourcePath+"css")
	r.Static("/src", resourcePath+"src")
	r.LoadHTMLGlob(resourcePath + "html/*")
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	r.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", gin.H{})
	})
	r.POST("/play", func(c *gin.Context) {
		c.HTML(http.StatusOK, "play.html", gin.H{})
	})
	r.GET("/play", func(c *gin.Context) {
		c.HTML(http.StatusOK, "play.html", gin.H{})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	port := strconv.Itoa(config.ServiceConf.ListenPort)
	r.Run(":" + port)
}

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

	// Contact the server and print out its response.
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
	c.Redirect(http.StatusTemporaryRedirect, "/play")
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

	// Contact the server and print out its response.
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
	c.Redirect(http.StatusTemporaryRedirect, "/play")
}
