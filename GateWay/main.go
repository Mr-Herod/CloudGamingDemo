package main

import (
	"net/http"
	"strconv"

	"github.com/Mr-Herod/CloudGamingDemo/Gateway/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	resourcePath := config.ServiceConf.ResourcePath
	// 1.创建路由
	r := gin.Default()
	r.POST("/sendDes", HandleDes)
	r.Static("/js", resourcePath+"js")
	r.Static("/css", resourcePath+"css")
	r.Static("/src", resourcePath+"src")
	r.LoadHTMLGlob(resourcePath + "html/*")
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.html", gin.H{})
	})
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
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
