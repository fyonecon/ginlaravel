package config

// 服务配置
func GetServerConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)

	serverConfig["HOST"] = "127.0.0.1"            // 监听地址
	serverConfig["PORT"] = "8082"               // 监听端口
	serverConfig["VIEWS_PATTERN"] = "views/*/*" // 模板路径pattern
	serverConfig["ENV"] = "debug"               // 环境模式 release/debug/test

	return
}
