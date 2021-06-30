package Common
// 超全局变量

// globalData 全局可变参数
// 预先定义好键和值最好
var globalData map[string]interface{} = map[string]interface{}{
	"test": "0",
	"cpu_num": 0,
	"cpu_percent": 0.00,
	"cpu_load1": 0.00,
	"cpu_load5": 0.00,
	"cpu_load15": 0.00,

}

// SetGlobalData 设置或新增全局变量
func SetGlobalData(key string, value interface{}) {
	globalData[key] = value
}

// GetGlobalData 获取全局变量
func GetGlobalData(key string) interface{} {
	return globalData[key]
}

// DelGlobalData 删除全局变量
func DelGlobalData(key string){
	delete(globalData, key)
}