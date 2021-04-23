package config

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

// MySQL数据库配置
func GetMySQLConfig() map[string]string {
	// 初始化数据库配置map
	MySQLConfig := make(map[string]string)

	MySQLConfig["DB_HOST"] = "192.168.131.7"
	MySQLConfig["DB_PORT"] = "3306"
	MySQLConfig["DB_NAME"] = "ginlaravel"
	MySQLConfig["DB_USER"] = "root2"
	MySQLConfig["DB_PWD"] = "123456"
	MySQLConfig["DB_CHARSET"] = "utf8mb4"

	MySQLConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	MySQLConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	MySQLConfig["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return MySQLConfig
}
