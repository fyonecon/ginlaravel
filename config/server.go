package config

// GetServerConfig http服务配置
func GetServerConfig() map[string]string {
	conf := make(map[string]string)

	conf["HOST"] = "127.0.0.1"         // 监听地址
	conf["PORT"] = "8090"              // 监听端口
	conf["ENV"] = "release"            // 环境模式 release/debug/test

	return conf
}
