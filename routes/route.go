package routes

import (
	"ginlaravel/app/http/Controller"
	"ginlaravel/app/http/Controller/Gen1Controller"
	"ginlaravel/app/http/Controller/TestController"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	// ==默认==
	route.Any("/", func (ctx *gin.Context) {
		ctx.String(404, "狗子，空的路由地址会直接返回404。Gin4Laravel、Go4Laravel。")
	})

	// ==测试==
	test := route.Group("/test/")
	{ // 按分组注册路由
		test.Any("", Controller.Null)       // 空路由
		test.Any("tpl", TestController.Tpl) // 模版输出
		test.Any("api", TestController.Api) // 接口输出-简单数据
		test.Any("api2", TestController.Api2) // 直接接口输出-复杂数据
	}

	// ==版本1的接口分组==
	gen1 := route.Group("/gen1/")
	{
		gen1.Any("", Controller.Null) // 空路由

		gen1.Any("app/list_user", Gen1Controller.ListUser)
		gen1.Any("app/that_user", Gen1Controller.ThatUser)
		gen1.Any("app/add_user", Gen1Controller.AddUser)
		gen1.Any("app/update_user", Gen1Controller.UpdateUser)
		gen1.Any("app/del_user", Gen1Controller.DelUser)
		gen1.Any("app/clear_user", Gen1Controller.ClearUser)

	}



}
