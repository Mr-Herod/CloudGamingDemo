package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Mr-Herod/CloudGamingDemo/Gateway/config"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Llongfile) // set flags
	config.Init()
	resourcePath := config.ServiceConf.ResourcePath
	// 1.创建路由
	r := gin.Default()
	// 2.导入Html模板文件
	r.LoadHTMLGlob(resourcePath + "html/*")
	r.Static("/img", resourcePath+"img")
	r.Static("/js", resourcePath+"js")
	r.Static("/css", resourcePath+"css")
	r.Static("/src", resourcePath+"src")
	// 3.绑定路由规则
	r.POST("/startGame", HandleStartGame)
	r.POST("/signup", HandleSignUp)
	r.POST("/signin", HandleSignIn)
	r.POST("/getRank", HandleGetRank)
	r.POST("/rank", HandleGetRank)
	r.POST("/center", func(c *gin.Context) { c.HTML(http.StatusOK, "center.html", gin.H{}) })
	r.POST("/play", func(c *gin.Context) {
		c.SetCookie("gamename", "推箱子", 60*60*24, "/", "mrherod.cn", false, true)
		c.HTML(http.StatusOK, "play.html", gin.H{})
	})

	r.GET("/rank", HandleGetRank)
	r.GET("/signup", func(c *gin.Context) { c.HTML(http.StatusOK, "signup.html", gin.H{}) })
	r.GET("/center", func(c *gin.Context) { c.HTML(http.StatusOK, "center.html", gin.H{}) })
	r.GET("/signin", func(c *gin.Context) { c.HTML(http.StatusOK, "signin.html", gin.H{}) })
	r.GET("/play", func(c *gin.Context) {
		c.SetCookie("gamename", "推箱子", 60*60*24, "/", "mrherod.cn", false, true)
		c.HTML(http.StatusOK, "play.html", gin.H{})
	})
	r.POST("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", gin.H{}) })
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", gin.H{}) })
	// 4.启动服务
	r.Run(":" + strconv.Itoa(config.ServiceConf.ListenPort))
}
