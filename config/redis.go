package config

// GetRedisConfig Redis数据库配置
func GetRedisConfig() map[string]string {
	conf := make(map[string]string)

	conf["Addr"] = "127.0.0.1:6379"
	conf["Password"] = "" // no password set
	conf["DB"] = "0" // use default DB

	return conf
}
