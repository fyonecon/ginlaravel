package config

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

// GetFrameworkConfig 框架参数配置
func GetFrameworkConfig() map[string]string {

	// main.go文件的绝对路径
	conf := make(map[string]string)
	mainDirectory, _ := os.Getwd()
	mainDirectory = mainDirectory + "/"

	// 解读外部json配置
	type FrameworkJson struct {
		Timezone string `json:"timezone"`
		GlVersion string `json:"gl_version"`
		FrameworkDesc struct {
			Desc string `json:"desc"`
			Author string `json:"author"`
			Email string `json:"Email"`
			Github string `json:"github"`
		} `json:"framework_desc"`
	}
	frameworkJson, _ := os.Open(mainDirectory + "framework.json")
	var framework FrameworkJson
	defer frameworkJson.Close()
	decoder := json.NewDecoder(frameworkJson)
	err := decoder.Decode(&framework)
	if err != nil {
		log.Println("'framework.json' Decoder Error = ", err.Error())
		//os.Exit(200)
	}else {
		//log.Println(framework.FrameworkDesc.Desc)
	}
	if len(framework.Timezone) == 0 {
		framework.Timezone = "Asia/Shanghai"
	}
	if len(framework.GlVersion) == 0 {
		framework.GlVersion = "gl_version.default.0.0.0"
	}

	conf["timezone"] = framework.Timezone 		// 时区
	conf["gl_version"] = framework.GlVersion 	// GinLaravel（或Ginvel）版本信息

	conf["go_version"] = runtime.Version()		// go版本
	conf["go_root"] = runtime.GOROOT()

	conf["framework_path"] = mainDirectory 		// 默认使用框架根目录
	conf["storage_path"] = mainDirectory + "storage/" // 文件存储文件夹

	return conf
}