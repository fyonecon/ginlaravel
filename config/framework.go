package config

import (
	"os"
	"runtime"
)

// GetFrameworkConfig http服务配置
func GetFrameworkConfig() map[string]string {
	conf := make(map[string]string)
	mainDirectory, _ := os.Getwd() // main.go文件的绝对路径
	mainDirectory = mainDirectory + "/"

	conf["timezone"] = "Asia/Shanghai" 		// 时区
	conf["gl_version"] = "gl-1.8.21.0528.20" 	// GinLaravel版本信息
	conf["go_version"] = runtime.Version()
	conf["go_root"] = runtime.GOROOT()
	conf["go_path"] = mainDirectory // 默认使用框架gopath
	conf["storage_path"] = mainDirectory + "storage/" // 文件存储文件夹

	return conf
}