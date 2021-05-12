package driver
// 使用插件：https://github.com/go-redis/redis

import (
	"context"
	"ginlaravel/app/Common"
	"ginlaravel/config"
	"github.com/go-redis/redis/v8"
	"log"
)

var rdbConfig map[string]string = config.GetRedisConfig()
var RedisDb *redis.Client

func init() {
	log.Println("尝试连接GoRedis...")

	RedisDb = redis.NewClient(&redis.Options{ // 连接服务
		Addr:     rdbConfig["Addr"],                        // string
		Password: rdbConfig["Password"],                    // string
		DB: int(Common.StringToInt(rdbConfig["DB"])), 		// int
	})
	RedisPong, RedisErr := RedisDb.Ping(context.Background()).Result() // 心跳
	if RedisErr != nil {
		log.Println("Redis服务未运行。。。")
		log.Println(RedisPong, RedisErr)
		//os.Exit(200)
	}else {
		log.Println("GoRedis已连接 >>> ")
	}
}