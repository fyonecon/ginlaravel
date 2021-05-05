package Middleware
// 利用redis缓存来定期存储耗时任务的结果

import (
	"fmt"
	"ginlaravel/app/Common"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
	"time"
)

// GetCacheInput 查询缓存
func GetCacheInput(ctx *gin.Context, key string) map[string]interface{} {
	back, err := Kit.RDB.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		//back = fmt.Sprintf("%s", err)
		back = ""
	}
	return Common.StringToMap(back)
}

// CreateCacheInput 创建缓存
func CreateCacheInput(ctx *gin.Context, key string, back []byte) interface{} {
	err := Kit.RDB.Set(ctx, key, back, 0).Err()
	if err != nil {
		fmt.Println(err)
		return 0
	}else {
		// 设置键过期时间，s
		timeout := 1200*time.Second // s，默认每20min可更新一次
		res := Kit.RDB.Expire(ctx, key, timeout)
		fmt.Println(res)
		return res
		//return 1
	}
}