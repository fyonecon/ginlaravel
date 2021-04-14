package common


// 代码参数配置
var Config = map[string]string{
	"test": "test-config",
	"api": "http",
}


// 代码公共函数
func Test(txt string) string {
	if len(txt) <= 0 {
		txt = "txt-nil"
	}
	Config["test"] = txt
	return txt
}