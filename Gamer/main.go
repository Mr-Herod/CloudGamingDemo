package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mr-Herod/CloudGamingDemo/Gamer/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	InitGame()
	r := gin.Default()
	r.POST("/", handle)
	port := strconv.Itoa(config.ServiceConf.ListenPort)
	r.Run(":" + port)
}

func handle(c *gin.Context) {
	des := c.PostForm("des")
	fmt.Println("client des:", des[:10])
	desChan := make(chan string)
	go RTC(des, desChan)
	remoteDes := <-desChan
	c.String(http.StatusOK, remoteDes)
}
