package RouterGroup

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Http/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiGen1(route *gin.Engine){
	// api分组路由
	api := route.Group("/api/", Middleware.HttpCorsApi)
	{
		//
		gen1 := api.Group("/gen1/")
		gen1.Any("1", Middleware.HttpLimiter(2), func(ctx *gin.Context) {
			ctx.JSONP(http.StatusNotFound, gin.H{
				"state": 200,
				"msg": "gin1",
				"content": map[string]interface{}{
					"time": Common.GetTimeDate("Ymd.His.ms.ns"),
				},
			})
		})

	}


}
