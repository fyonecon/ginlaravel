package Gen1App

import (
	"fmt"
	"ginlaravel/app/Common"
	"ginlaravel/bootstrap/driver"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisDb *redis.Client = driver.RedisDb

// redis无序不重复
func RedisSet(ctx *gin.Context){

	err := redisDb.Set(ctx, "date", Common.GetTimeDate("Y-m-d H:i:s"), 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := redisDb.Get(ctx, "date").Result()
	if err != nil {
		val = fmt.Sprintf("%s", err)
	}

	// 接口返回
	ctx.JSON(200, gin.H{
		"state": 1,
		"msg": "GORedis示例",
		"content": map[string]string{
			"date": val,
		},
	})
}

// redis有序队列
func RedisList(ctx *gin.Context)  {

	// 按序写入数据
	for i:=0; i<3; i++ {
		_i := Common.IntToString(i)
		err := redisDb.LPush(ctx, "list_name_" + _i, Common.GetTimeDate("Y-m-d H:i:s")).Err()
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 1) // ns-ms
	}

	// 按序读出数据
	_len, _:= redisDb.LLen(ctx, "list_name_1").Result() // 获取表中元素的个数
	_val, _err := redisDb.LRange(ctx, "list_name_1", 0, _len-1).Result()
	if _err != nil {
		fmt.Println(_err)
	}
	fmt.Println(_val)

	// 接口返回
	ctx.JSON(200, gin.H{
		"state": 1,
		"msg": "GORedis示例",
		"content": [...]string{Common.IntToString(int(_len)), _val[0], _val[1], _val[2], _val[3]},
	})
}