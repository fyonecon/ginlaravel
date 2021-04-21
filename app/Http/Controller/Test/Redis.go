package Test

import (
	"fmt"
	"ginlaravel/app/Common"
	"ginlaravel/bootstrap/driver"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisDb *redis.Client = driver.RedisDb

func GoRedis(ctx *gin.Context){

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