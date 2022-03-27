package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	common "github.com/Mr-Herod/CloudGamingDemo/Common"
)

type ServiceConfig struct {
	ServiceName string `json:"serviceName"`
	ListenPort  int    `json:"listenPort"`
}

var ServiceConf ServiceConfig

func Init() {
	file, err := os.Open(common.GetCurrentDirectory() + "/config/config.json")
	if err != nil {
		panic(fmt.Sprintf("open config error,err:%v", err))
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	err = json.Unmarshal(content, &ServiceConf)
	if err != nil {
		panic(fmt.Sprintf("config json.Unmarshal error,err:%v", err))
	}
	log.Printf("ServiceConf: %+v", ServiceConf)
}
