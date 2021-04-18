package routes

/*
	第二作者Author：fyonecon
	博客Blog：https://blog.csdn.net/weixin_41827162
	邮箱Email：2652335796@qq.com，ikydee@yahoo.com
	微信WeChat：fy66881159
	所在城市City：长沙ChangSha
	==2021年04月17号==
	==2016-2020#WEB、2018-2020#PHP、2021—#GO==
*/

import (
	"ginlaravel/app/http/Controller"
	"ginlaravel/app/http/Controller/Gen1Controller"
	"ginlaravel/app/http/Controller/Gen2Controller"
	"ginlaravel/app/http/Controller/TestController"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"time"
)

// 路由周期：请求路由名——>拦截请求频率——>校验请求方法和参数——>运行目标函数
func RegisterRoutes(route *gin.Engine) {

	// 拦截http请求频率
	var tbOptions  limiter.ExpirableOptions
	tbOptions.DefaultExpirationTTL = time.Second // 默认按每秒
	lmt := tollbooth.NewLimiter(3, &tbOptions) // 默认3次/秒，建议范围[1，20]

	// ==默认路由==
	route.Any("/", tollbooth_gin.LimitHandler(lmt), func (ctx *gin.Context) {
		ctx.String(404, "狗子，空的路由地址会直接返回404")
	})

	// ==测试==
	test := route.Group("/test/", tollbooth_gin.LimitHandler(lmt), Controller.VerifyTest)
	{ // 按分组注册路由
		test.Any("", Controller.Null)             // 空路由
		test.Any("test1", TestController.Test1)   // 空路由
		test.Any("tpl", TestController.Tpl)       // 模版输出
		test.Any("api", TestController.Api)       // 接口输出-简单数据
		test.Any("api2", TestController.Api2)     // 直接接口输出-复杂数据
		test.Any("init", TestController.Test2Run) // 直接接口输出-复杂数据
	}

	// ==版本1的接口分组==
	gen1 := route.Group("/gen1/", tollbooth_gin.LimitHandler(lmt), Controller.VerifyGen1)
	{
		gen1.Any("", tollbooth_gin.LimitHandler(lmt), Controller.Null) // 空路由

		gen1.Any("app/list_user", tollbooth_gin.LimitHandler(lmt), Gen1Controller.ListUser)
		gen1.Any("app/that_user", tollbooth_gin.LimitHandler(lmt), Gen1Controller.ThatUser)
		gen1.Any("app/add_user", tollbooth_gin.LimitHandler(lmt), Gen1Controller.AddUser)
		gen1.Any("app/update_user", tollbooth_gin.LimitHandler(lmt), Gen1Controller.UpdateUser)
		gen1.Any("app/del_user", tollbooth_gin.LimitHandler(lmt), Gen1Controller.DelUser)
		gen1.Any("app/clear_user", tollbooth_gin.LimitHandler(lmt), Gen1Controller.ClearUser)

	}

	// ==版本2的接口分组==
	gen2 := route.Group("/gen2/", tollbooth_gin.LimitHandler(lmt), Controller.VerifyGen2)
	{
		gen2.Any("", Controller.Null) // 空路由

		gen2.Any("app/list_user", Gen2Controller.ListUser)
		gen2.Any("app/that_user", Gen2Controller.ThatUser)
		//gen2.Any("app/add_user", Gen2Controller.AddUser)
		//gen2.Any("app/update_user", Gen2Controller.UpdateUser)
		//gen2.Any("app/del_user", Gen2Controller.DelUser)
		//gen2.Any("app/clear_user", Gen2Controller.ClearUser)

	}



}
