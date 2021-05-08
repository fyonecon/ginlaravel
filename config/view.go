package config

// GetViewConfig html模版视图路径配置
func GetViewConfig() map[string]string {
	conf := make(map[string]string)

	// html模板文件路径。**代表文件夹，*代表文件。*结尾。
	conf["View_Pattern"] = "views/html/**/**/*"
	// 多静态文件的主文件夹。/结尾。
	conf["View_Static"] = "views/static/"

	return conf
}
