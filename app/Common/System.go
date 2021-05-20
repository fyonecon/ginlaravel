package Common
// 项目函数运行所依赖的参数配置

import (
	"os"
	"runtime"
)

// Config 代码参数配置
var Config = map[string]interface{}{
	"debug": true,
	"api": "http",
}

// ServerInfo 服务器信息
var ServerInfo = map[string]string{
	"timezone": "Asia/Shanghai", 		// 时区
	"gl_version": "gl-1.7.21.0520.16", 	// GinLaravel版本信息
	"go_version": runtime.Version(),
	"root_path": runtime.GOROOT(),
	"go_path": os.Getenv("GOPATH") + "/src/ginlaravel/",
	"storage_path": os.Getenv("GOPATH") + "/src/ginlaravel/storage/", // 文件存储在服务器的地址
}

// Page 分页参数
var Page = map[string]int{
	"limit": 20, // 每页多少条数据
	"max_page": 500, // 最大查到多少页
	"min_limit": 2, // 最少一次查多少条数据
	"max_limit": 100, // 最多一次查多少条数据
}
