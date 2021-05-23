package Kit
// 利用redis-RAM缓存来定期存储耗时任务的结果

import (
	"encoding/json"
	"fmt"
	"ginvel.com/app/Common"
	"github.com/gin-gonic/gin"
	"time"
)

// CreateCacheInput 创建缓存
func CreateCacheInput(ctx *gin.Context, key string, back map[string]interface{}) interface{} {
	backJson, _ := json.Marshal(back)
	err := RDB.Set(ctx, key, backJson, 0).Err()
	if err != nil {
		fmt.Println(err)
		return 0
	}else {
		// 设置键过期时间，s
		timeout := 1200*1000*time.Millisecond // ms，默认每20min可更新一次
		res := RDB.PExpire(ctx, key, timeout)
		fmt.Println(res)
		return res
		//return 1
	}
}

// GetCacheInput 查询缓存
func GetCacheInput(ctx *gin.Context, key string) map[string]interface{} {
	back, err := RDB.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		//back = fmt.Sprintf("%s", err)
		back = ""
	}
	return Common.JsonStringToMap(back)
}