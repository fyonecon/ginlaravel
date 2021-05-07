package config

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

// GetServerConfig http服务配置
func GetServerConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)

	serverConfig["HOST"] = "127.0.0.1"         // 监听地址
	serverConfig["PORT"] = "8090"              // 监听端口
	serverConfig["VIEWS_PATTERN"] = "views/pages/**/*" // html模板文件路径。**代表文件夹，*代表文件
	serverConfig["ENV"] = "release"            // 环境模式 release/debug/test

	return
}
