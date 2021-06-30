package config

import (
	"encoding/json"
	"log"
	"os"
)

// GetMySQLConfig MySQL数据库配置
func GetMySQLConfig() map[string]string {
	// FrameworkJson 解读外部json配置
	// 优点：可以不更改二进制文件的情况下更改参数
	type FrameworkJson struct {
		Timezone string `json:"timezone"`
		Config struct {
			Mysql struct {
				HOST string `json:"DB_HOST"`
				PORT string `json:"DB_PORT"`
				NAME string `json:"DB_NAME"`
				USER string `json:"DB_USER"`
				PWD string `json:"DB_PWD"`
				CHARSET string `json:"DB_CHARSET"`
			} `json:"mysql"`
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

	host := framework.Config.Mysql.HOST
	port := framework.Config.Mysql.PORT
	name := framework.Config.Mysql.NAME
	user := framework.Config.Mysql.USER
	pwd := framework.Config.Mysql.PWD
	charset := framework.Config.Mysql.CHARSET

	// 默认值
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "3306"
	}
	if len(name) == 0 || len(user) == 0 {
		log.Println("mysql数据库名或账户名不能为空。。。")
		os.Exit(200)
	}
	if len(charset) == 0 {
		charset = "utf8mb4"
	}

	conf := make(map[string]string)

	conf["DB_HOST"] = host // 127.0.0.1
	conf["DB_PORT"] = port
	conf["DB_NAME"] = name
	conf["DB_USER"] = user
	conf["DB_PWD"] = pwd
	conf["DB_CHARSET"] = charset // "utf8mb4"
	conf["DB_TIMEOUT"] = "12s"

	conf["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	conf["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	conf["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return conf
}
