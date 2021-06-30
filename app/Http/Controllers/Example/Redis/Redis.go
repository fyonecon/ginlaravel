package Redis

import (
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
	"time"
)

// RedisSet redis无序不重复
func RedisSet(ctx *gin.Context){

	err := Kit.RDB.Set(ctx, "date", Common.GetTimeDate("Y-m-d H:i:s"), 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := Kit.RDB.Get(ctx, "date").Result()
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

// RedisList redis有序队列
func RedisList(ctx *gin.Context)  {

	// 按序写入数据
	for i:=0; i<3; i++ {
		_i := Common.IntToString(int64(i))
		err := Kit.RDB.LPush(ctx, "list_name_" + _i, Common.GetTimeDate("Y-m-d H:i:s")).Err()
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 1) // ns-ms
	}

	// 按序读出数据
	_len, _:= Kit.RDB.LLen(ctx, "list_name_1").Result() // 获取表中元素的个数
	_val, _err := Kit.RDB.LRange(ctx, "list_name_1", 0, _len-1).Result()
	if _err != nil {
		fmt.Println(_err)
	}
	fmt.Println(_val)

	// 接口返回
	ctx.JSON(200, gin.H{
		"state": 1,
		"msg": "GORedis示例",
		"content": [...]string{Common.IntToString(_len), _val[0], _val[1], _val[2], _val[3]},
	})
}