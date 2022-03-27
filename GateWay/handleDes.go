package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Des struct {
	Des string `json:"des"`
}

func HandleDes(c *gin.Context) {
	var json Des
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rspStr, err := remoteHandle(json.Des)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, rspStr)
}

func remoteHandle(des string) (string, error) {
	data := url.Values{
		"des": {des},
	}
	resp, err := http.PostForm("http://localhost:8010", data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	rsp := string(body)
	return rsp, nil
}
