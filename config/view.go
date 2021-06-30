package config

import (
	"encoding/json"
	"log"
	"os"
)

// GetViewConfig html模版视图路径配置
func GetViewConfig() map[string]string {
	// FrameworkJson 解读外部json配置
	// 优点：可以不更改二进制文件的情况下更改参数
	type FrameworkJson struct {
		Timezone string `json:"timezone"`
		Config struct {
			View struct {
				ViewPattern string `json:"View_Pattern"`
				ViewStatic string `json:"View_Static"`
			} `json:"view"`
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

	pattern := framework.Config.View.ViewPattern
	static := framework.Config.View.ViewStatic

	// 默认值
	if len(pattern) == 0 {
		pattern = "views/html/**/**/*"
	}
	if len(static) == 0 {
		static = "views/static/"
	}

	conf := make(map[string]string)

	// html模板文件路径。**代表文件夹，*代表文件。*结尾。
	conf["View_Pattern"] = pattern
	// 多静态文件的主文件夹。/结尾。
	conf["View_Static"] = static

	return conf
}
