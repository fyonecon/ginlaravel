package config

import (
	"encoding/json"
	"log"
	"os"
)

// GetServerConfig http服务配置
func GetServerConfig() map[string]string {
	// FrameworkJson 解读外部json配置
	// 优点：可以不更改二进制文件的情况下更改参数
	type FrameworkJson struct {
		Timezone string `json:"timezone"`
		Config struct {
			Server struct {
				HOST string `json:"HOST"`
				PORT string `json:"PORT"`
				ENV string `json:"ENV"`
			} `json:"server"`
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

	host := framework.Config.Server.HOST
	port := framework.Config.Server.PORT
	env := framework.Config.Server.ENV

	// 默认值
	// docker中运行请使用：0.0.0.0，本地测试请使用：127.0.0.1
	if host == "localhost" || len(host) == 0{
		host = "0.0.0.0"
	}
	if len(port) == 0 {
		port = "8090"
	}
	if len(env) == 0 {
		env = "release"
	}

	conf := make(map[string]string)

	conf["HOST"] = host // 监听地址，部署在docker中请使用：0.0.0.0。建议不要用127.0.0.1或localhost
	conf["PORT"] = port // 监听端口
	conf["ENV"] = env // 环境模式 release/debug/test

	return conf
}
