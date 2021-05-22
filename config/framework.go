package config

import (
	"log"
	"os"
	"runtime"
)

// GetFrameworkConfig http服务配置
func GetFrameworkConfig() map[string]string {
	conf := make(map[string]string)

	// 系统gopath
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = "/home/wwwroot/go" // centos获取不到env？？？手动填吧
		log.Println("env中GOPATH获取为空，将使用默认手动地址（/config/framework.go:17行）。默认路径：", gopath)
	}
	// 框架gopath
	frameworkPath := gopath + "/src/ginlaravel/"

	conf["timezone"] = "Asia/Shanghai" 		// 时区
	conf["gl_version"] = "gl-1.7.21.0522.15" 	// GinLaravel版本信息
	conf["go_version"] = runtime.Version()
	conf["go_root"] = runtime.GOROOT()
	conf["go_path"] = frameworkPath // 默认使用框架gopath
	conf["storage_path"] = frameworkPath + "storage/" // 文件存储在服务器的地址

	return conf
}