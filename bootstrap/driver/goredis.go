package driver
// 使用插件：https://github.com/go-redis/redis

import (
	"context"
	"ginvel.com/app/Common"
	"ginvel.com/config"
	"github.com/go-redis/redis/v8"
	"log"
)

var RedisDb *redis.Client

func InitRedis() {
	log.Println("尝试连接GoRedis...")

	var rdbConfig map[string]string = config.GetRedisConfig()

	RedisDb = redis.NewClient(&redis.Options{ // 连接服务
		Addr:     rdbConfig["Addr"],                        // string
		Password: rdbConfig["Password"],                    // string
		DB: int(Common.StringToInt(rdbConfig["DB"])), 		// int
	})
	RedisPong, RedisErr := RedisDb.Ping(context.Background()).Result() // 心跳
	if RedisErr != nil {
		log.Println("Redis服务未运行。。。", RedisPong, RedisErr)
		log.Println("Redis常用命令：\n" +
			" 启动：src/redis-server \n" +
			" 进入命令行：src/redis-cli \n" +
			" 关闭安全模式：CONFIG SET protected-mode no \n" +
			" 重置密码：config set requirepass [密码]\n")
		//os.Exit(200)
	}else {
		log.Println("GoRedis已连接 >>> ")
	}
}