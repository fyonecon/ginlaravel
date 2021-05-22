package config

import (
	"os"
	"runtime"
)

// GetFrameworkConfig http服务配置
func GetFrameworkConfig() map[string]string {
	conf := make(map[string]string)
	frameworkPath := os.Getenv("GOPATH") + "/src/ginlaravel/"

	conf["timezone"] = "Asia/Shanghai" 		// 时区
	conf["gl_version"] = "gl-1.7.21.0522.15" 	// GinLaravel版本信息
	conf["go_version"] = runtime.Version()
	conf["root_path"] = runtime.GOROOT()
	conf["go_path"] = frameworkPath
	conf["storage_path"] = frameworkPath + "storage/" // 文件存储在服务器的地址

	return conf
}