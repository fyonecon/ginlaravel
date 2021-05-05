package config

// GetRedisConfig Redis数据库配置
func GetRedisConfig() map[string]string {
	rdbConfig := make(map[string]string)

	rdbConfig["Addr"] = "127.0.0.1:6379"
	rdbConfig["Password"] = "" // no password set
	rdbConfig["DB"] = "0" // use default DB

	return rdbConfig
}
