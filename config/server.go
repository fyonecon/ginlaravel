package config

// GetServerConfig http服务配置
func GetServerConfig() map[string]string {
	conf := make(map[string]string)

	conf["HOST"] = "0.0.0.0"           // 监听地址，部署在docker中请使用：0.0.0.0
	conf["PORT"] = "8090"              // 监听端口
	conf["ENV"] = "release"            // 环境模式 release/debug/test

	return conf
}
