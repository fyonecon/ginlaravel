package config

// MySQL数据库配置
func GetDbConfig() map[string]string {
	// 初始化数据库配置map
	dbConfig := make(map[string]string)

	dbConfig["DB_HOST"] = "192.168.131.7"
	dbConfig["DB_PORT"] = "3306"
	dbConfig["DB_NAME"] = "ginlaravel"
	dbConfig["DB_USER"] = "root2"
	dbConfig["DB_PWD"] = "123456"
	dbConfig["DB_CHARSET"] = "utf8mb4"

	dbConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	dbConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	dbConfig["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return dbConfig
}
