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
	"ginlaravel/app/Http/Controller/Gen1"
	"ginlaravel/app/Http/Controller/Gen1/Gen1App"
	"ginlaravel/app/Http/Controller/Gen1/Gen1User"
	"ginlaravel/app/Http/Controller/Gen3"
	"ginlaravel/app/Http/Controller/Gen3/Gen3App"
	"ginlaravel/app/Http/Controller/Gen3/Gen3Open"
	_ "ginlaravel/app/Http/Controller/Gen3/Gen3User"
	"ginlaravel/app/Http/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由周期：请求路由名——>拦截请求频率——>header——>校验请求方法和Token参数——>运行目标函数
func RegisterRoutes(route *gin.Engine) {
	// ==系统必要路由==
	route.NoRoute(Middleware.HttpCors(), Middleware.HttpLimiter(2), func (ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"state": 404, "msg": "未定义此名称的路由名", "content": ctx.Request.URL.Path,
		})
	})
	route.Any("/",  Middleware.HttpCors(), Middleware.HttpLimiter(2), func (ctx *gin.Context) {
		//ctx.String(http.StatusNotFound, "请先指定路由（404），" + ctx.Request.Method + "。")
		ctx.JSON(http.StatusNotFound, gin.H{
			"state": 403, "msg": "指定路由名后才可访问", "content": ctx.Request.Method,
		})
	})

	// ==版本1的接口分组==
	gen1 := route.Group("/gen1/", Middleware.HttpCors())
	{
		// 访问：域名/gen1/user/xxx
		user := gen1.Group("/user/", Middleware.HttpLimiter(2), Gen1.VerifyGen1User)
		{
			user.Any("list_user", Gen1User.ListUser)
			user.Any("that_user", Gen1User.ThatUser)
			user.Any("add_user", Gen1User.AddUser)
			user.Any("update_user", Gen1User.UpdateUser)
			user.Any("del_user", Gen1User.DelUser)
			user.Any("clear_user", Gen1User.ClearUser)
		}

		// 访问：域名/gen1/app/xxx
		app := gen1.Group("/app/", Middleware.HttpLimiter(2), Gen1.VerifyGen1App)
		{
			app.Any("test1", Gen1App.Test1)                      // 空路由
			app.Any("tpl", Gen1App.Tpl)                          // 模版输出
			app.Any("api", Gen1App.Api)                          // 接口输出-简单数据
			app.Any("api2", Gen1App.Api2)                        // 直接接口输出-复杂数据
			app.Any("init", Gen1App.Test2Run)                    // 直接接口输出-复杂数据
			app.Any("upload_form_file", Gen3Open.UploadFormFile) // 直接接口输出-复杂数据
			app.Any("redis_set", Gen1App.RedisSet)
			app.Any("redis_list", Gen1App.RedisList)
		}

	}

	// ==版本3的接口分组==
	gen3 := route.Group("/gen3/", Middleware.HttpCors())
	{
		gen3.Any("app/get_app_token", Middleware.HttpLimiter(2), Gen3.VerifyOpen, Gen3App.GetAppToken)
		gen3.Any("user/user_login", Middleware.HttpLimiter(2), Gen3.VerifyOpen, Gen3App.GetAppToken)

		// 访问：域名/gen3/open/xxx
		open := gen3.Group("/open/", Middleware.HttpLimiter(2), Gen3.VerifyOpen)
		{
			open.Any("upload_form_file", Gen3Open.UploadFormFile)
		}

		//// 访问：域名/gen3/app/xxx
		//app := gen3.Group("/user/", Middleware.HttpLimiter(4), Gen3.VerifyApp)
		//{
		//	app.Any("list_user", Gen3User.ListUser)
		//	//app.Any("list_user", Gen3User.thatUser)
		//}
		//
		//// 访问：域名/gen3/user/xxx
		//user := gen3.Group("/user/", Middleware.HttpLimiter(4), Gen3.VerifyUser)
		//{
		//	user.Any("list_user", Gen3User.ListUser)
		//	//user.Any("list_user", Gen3User.thatUser)
		//}

	}



}
