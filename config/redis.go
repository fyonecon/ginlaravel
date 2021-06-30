package config

import (
	"encoding/json"
	"log"
	"os"
)

// GetRedisConfig Redis数据库配置
func GetRedisConfig() map[string]string {
	// FrameworkJson 解读外部json配置
	// 优点：可以不更改二进制文件的情况下更改参数
	type FrameworkJson struct {
		Timezone string `json:"timezone"`
		Config struct {
			Redis struct {
				Addr string `json:"Addr"`
				Password string `json:"Password"`
				DB string `json:"DB"`
			} `json:"redis"`
		} `json:"config"`
	}

	// main.go文件的绝对路径
	mainDirectory, _ := os.Getwd()
	mainDirectory = mainDirectory + "/"

	// 加载配置文件
	frameworkJson, _ := os.Open(mainDirectory + "framework.json")
	var framework FrameworkJson
	defer frameworkJson.Close()
	decoder := json.NewDecoder(frameworkJson)
	err := decoder.Decode(&framework)
	if err != nil {
		log.Println("'framework.json' Decoder Error = ", err.Error(), "运行中断")
		os.Exit(200)
	}

	// 默认值
	addr := framework.Config.Redis.Addr
	pwd := framework.Config.Redis.Password
	db := framework.Config.Redis.DB
	if len(addr) == 0 {
		addr = "127.0.0.1:6379"
	}
	if len(db) == 0 {
		db = "0"
	}

	conf := make(map[string]string)

	conf["Addr"] = addr // 例子：127.0.0.1:6379
	conf["Password"] = pwd // 无密码就设置为：""
	conf["DB"] = db // use default DB

	return conf
}
