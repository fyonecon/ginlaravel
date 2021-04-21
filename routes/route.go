package routes

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

import (
	"ginlaravel/app/Http/Controller/Gen2"
	"ginlaravel/app/Http/Controller/Gen3"
	"ginlaravel/app/Http/Controller/Gen3/Gen3App"
	"ginlaravel/app/Http/Controller/Gen3/Gen3Open"
	"ginlaravel/app/Http/Controller/Test"
	"ginlaravel/app/Http/Middleware"
	"ginlaravel/app/http/Controller"
	"ginlaravel/app/http/Controller/Gen1Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由周期：请求路由名——>拦截请求频率——>header——>校验请求方法和Token参数——>运行目标函数
func RegisterRoutes(route *gin.Engine) {
	// ==默认路由==
	route.Any("/", Middleware.HttpLimiter(1), Middleware.HttpCors(), func (ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "请先指定路由（404）")
	})
	route.NoRoute(Middleware.HttpLimiter(1), Middleware.HttpCors(), func (ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"state": 404, "msg": "未定义此名称的路由", "content": ctx.Request.URL,
		})
	})

	// ==测试==
	test := route.Group("/test/", Middleware.HttpLimiter(3), Middleware.HttpCors(), Controller.VerifyTest)
	{ // 按分组注册路由
		test.Any("test1", Test.Test1)                         // 空路由
		test.Any("tpl", Test.Tpl)                             // 模版输出
		test.Any("api", Test.Api)                             // 接口输出-简单数据
		test.Any("api2", Test.Api2)                           // 直接接口输出-复杂数据
		test.Any("init", Test.Test2Run)                       // 直接接口输出-复杂数据
		test.Any("upload_form_file", Gen3Open.UploadFormFile) // 直接接口输出-复杂数据
		test.Any("redis_set", Test.RedisSet)
		test.Any("redis_list", Test.RedisList)
	}

	// ==版本1的接口分组==
	gen1 := route.Group("/gen1/", Middleware.HttpLimiter(4), Middleware.HttpCors(), Controller.VerifyGen1)
	{
		gen1.Any("app/list_user", Gen1Controller.ListUser)
		gen1.Any("app/that_user", Gen1Controller.ThatUser)
		gen1.Any("app/add_user", Gen1Controller.AddUser)
		gen1.Any("app/update_user", Gen1Controller.UpdateUser)
		gen1.Any("app/del_user", Gen1Controller.DelUser)
		gen1.Any("app/clear_user", Gen1Controller.ClearUser)
	}

	// ==版本2的接口分组==
	gen2 := route.Group("/gen2/", Middleware.HttpLimiter(4), Middleware.HttpCors(), Controller.VerifyGen2)
	{
		gen2.Any("app/list_user", Gen2.ListUser)
		gen2.Any("app/that_user", Gen2.ThatUser)
		//gen2.Any("app/add_user", Gen2.AddUser)
		//gen2.Any("app/update_user", Gen2.UpdateUser)
		//gen2.Any("app/del_user", Gen2.DelUser)
		//gen2.Any("app/clear_user", Gen2.ClearUser)

	}

	// ==版本3的接口分组==
	route.Any("/gen3/app/get_app_token", Middleware.HttpLimiter(1), Middleware.HttpCors(), Gen3.VerifyOpen, Gen3App.GetAppToken)
	route.Any("/gen3/user/user_login", Middleware.HttpLimiter(1), Middleware.HttpCors(), Gen3.VerifyOpen, Gen3App.GetAppToken)
	route.Any("/gen3/open/upload_form_file", Middleware.HttpLimiter(1), Middleware.HttpCors(), Gen3.VerifyOpen, Gen3Open.UploadFormFile)

	gen3 := route.Group("/gen3/", Middleware.HttpLimiter(4), Middleware.HttpCors(), Gen3.VerifyApp)
	{
		gen3.Any("app/list_user", Gen3App.ListUser)
		//gen3.Any("app/list_user", Gen3App.thatUser)

	}



}
