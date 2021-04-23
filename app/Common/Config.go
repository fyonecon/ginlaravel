package Common
// 项目函数（或第三方服务）运行所依赖的参数配置

import (
	"os"
	"runtime"
)

// 代码参数配置
var Config = map[string]string{
	"test": "test-config",
	"api": "http",
}

// 服务器信息
var ServerInfo = map[string]string{
	"root_path": runtime.GOROOT(),
	"go_path": os.Getenv("GOPATH") + "/src/ginlaravel/",
	"storage_path": os.Getenv("GOPATH") + "/src/ginlaravel/storage/", // 文件存储在服务器的地址
}

// 分页参数
var Page = map[string]int{
	"limit": 20, // 每页多少条数据
	"max_page": 500, // 最大查到多少页
	"min_limit": 2, // 最少一次查多少条数据
	"max_limit": 100, // 最多一次查多少条数据
}

// 小程序数据
var XCX_WX_KY = map[string]string{
	"AppId": "xxx",
	"AppSecret": "xxx",
	"img_domain": "https://img.xxx.com",
}