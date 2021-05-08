package config

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

// GetMySQLConfig MySQL数据库配置
func GetMySQLConfig() map[string]string {
	// 初始化数据库配置map
	conf := make(map[string]string)

	conf["DB_HOST"] = "192.168.131.7"
	conf["DB_PORT"] = "3306"
	conf["DB_NAME"] = "ginlaravel"
	conf["DB_USER"] = "root2"
	conf["DB_PWD"] = "123456"
	conf["DB_CHARSET"] = "utf8mb4"
	conf["DB_TIMEOUT"] = "12s"

	conf["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	conf["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	conf["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return conf
}
