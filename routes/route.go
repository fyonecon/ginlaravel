package routes

import (
	"ginlaravel/app/http/Controller"
	"ginlaravel/app/http/Controller/Gen1"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	route.Any("/", func (ctx *gin.Context) {
		ctx.String(404, "狗子，空的路由地址会直接返回404。Gin4Laravel、Go4Laravel。")
	})                                  // 空路由


	gen1 := route.Group("/gen1/")
	{ // 按分组注册路由
		gen1.Any("", Controller.Null) // 空路由
		gen1.Any("tpl", Gen1.Tpl)      // 模版输出
		gen1.Any("api", Gen1.Api)      // 接口输出
		gen1.Any("api/db", Gen1.ApiDB) // 读数据接口（不依赖模型）输出
		gen1.Any("api/md", Gen1.ApiMD) // 读数据模型（model）接口输出

	}


	gen2 := route.Group("/gen2/")
	{ // 按分组注册路由
		gen2.Any("", Controller.Null) // 空路由

		gen2.Any("app", Controller.Null) // 空路由

		gen2.Any("user", Controller.Null) // 空路由

		gen2.Any("admin", Controller.Null) // 空路由

	}



}
