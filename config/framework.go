package config

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
)

// GetFrameworkConfig 框架参数配置
func GetFrameworkConfig() map[string]string {

	// FrameworkJson 解读外部json配置
	// 优点：可以不更改二进制文件的情况下更改参数
	type FrameworkJson struct {
		GlVersion string `json:"gl_version"`
		FrameworkDesc struct {
			Desc string `json:"Desc"`
			Author string `json:"Author"`
			Email string `json:"Email"`
			Github string `json:"Github"`
		} `json:"framework_desc"`
		FrameworkPath string `json:"framework_path"`
		Timezone string `json:"timezone"`
		Common struct{
			LogDebug string `json:"log_debug"`
		}
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
		log.Println("'framework.json' framework：Decoder Error = ", err.Error(), "framework参数未拿到（可选，跳过）")
		//os.Exit(200)
	}

	timezone := framework.Timezone
	glVersion := framework.GlVersion
	logDebug := framework.Common.LogDebug

	if len(timezone) == 0 {
		timezone = "Asia/Shanghai"
	}
	if len(glVersion) == 0 {
		glVersion = "gl_version.default.0.0.0"
	}
	if logDebug != "true" || logDebug != "false" {
		logDebug = "true"
	}

	conf := make(map[string]string)

	conf["timezone"] = timezone // 时区
	conf["gl_version"] = glVersion 	 // GinLaravel（或Ginvel）版本信息

	conf["go_version"] = runtime.Version()		 // go版本
	conf["go_root"] = runtime.GOROOT()

	conf["framework_path"] = mainDirectory 		 // 默认使用框架根目录
	conf["storage_path"] = mainDirectory + "storage/" // 文件存储文件夹

	conf["log_debug"] = logDebug

	return conf
}