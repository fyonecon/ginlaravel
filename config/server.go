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
func GetServerConfig() map[string]string {
	conf := make(map[string]string)

	conf["HOST"] = "127.0.0.1"         // 监听地址
	conf["PORT"] = "8090"              // 监听端口
	conf["ENV"] = "release"            // 环境模式 release/debug/test

	return conf
}
