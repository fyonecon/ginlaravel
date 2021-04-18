package Common

// 代码参数配置
var Config = map[string]string{
	"test": "test-config",
	"api": "http",
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