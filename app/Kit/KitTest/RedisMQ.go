package KitTest
// 利用Redis制作的适合中小系统的消息队列MQ
// 大型系统就不要用这里了，用kafka吧

import (
	"encoding/json"
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
)

type CallFunc func(key string, id1 int64, id2 int64) int64

// CreateRedisMQ 创建缓存延迟执行队列事件
// 调用本函数建议用多线程（互斥锁）
func CreateRedisMQ(ctx *gin.Context, key string, id1 int64, id2 int64, callFunc CallFunc) int64 {
	// 消息事件和消息参数写入Redis
	k := "gl_cache_mq" // list可重复
	_v := map[string]interface{}{
		"key": key,
		"id1": id1,
		"id2": id2,
		"callFunc": callFunc,
	}
	v, _ := json.Marshal(_v)
	err := Kit.RDB.LPush(ctx, k, v).Err()
	if err != nil {
		fmt.Println(err)
		Kit.Log("CreateRedisMQ缓存写入失败", "")
		return 0
	}
	return 1
}

// DoRedisMQ 读取Redis队列数据，并生成队列的目标回调事件
// 将一次性执行完所有队列MQ事件（可定时调用但不能单次调用）
func DoRedisMQ(ctx *gin.Context) {
	// 按序读出数据
	_k := "gl_cache_mq"
	// 获取表中元素的个数
	_len, _ := Kit.RDB.LLen(ctx, _k).Result()
	_val, _err := Kit.RDB.LRange(ctx, _k, 0, _len-1).Result()
	if _err != nil {
		fmt.Println(_err)
		return
	}
	fmt.Println(_val)
	for i := 0; i < int(_len); i++ {
		val := _val[1]
		back := Common.JsonStringToMap(val)
		_key := back["key"]
		_id1 := back["id1"]
		_id2 := back["id2"]
		_callFunc := back["callFunc"]

		key := Common.ValueInterfaceToString(_key)
		id1 := Common.ValueInterfaceToInt(_id1)
		id2 := Common.ValueInterfaceToInt(_id2)
		callFunc := Common.ValueInterfaceToString(_callFunc)

		_CallFunc(key, id1, id2, callFunc)

	}

}

// 回调函数
func _CallFunc(key string, id1 int64, id2 int64, callFunc string) int64 {
	switch callFunc {
		case "test":
			// test()
			break
		case "test1":
			// test1()
			break
		case "test2":
			// test2()
			break
		default:
			Kit.Log("_CallFunc无回调函数", "")
			break
	}

	return 1
}