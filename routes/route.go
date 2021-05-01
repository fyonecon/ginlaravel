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
	"ginlaravel/app/Common"
	"ginlaravel/app/Http/Controller/Gen1"
	"ginlaravel/app/Http/Controller/Gen1/Gen1App"
	"ginlaravel/app/Http/Controller/Gen1/Gen1User"
	"ginlaravel/app/Http/Controller/Gen3"
	"ginlaravel/app/Http/Controller/Gen3/Gen3App"
	"ginlaravel/app/Http/Controller/Gen3/Gen3Open"
	"ginlaravel/app/Http/Controller/Gen3/Gen3User"
	"ginlaravel/app/Http/Middleware"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由访问原则：宽进严出，所以都用Any，在拦截器里面拦截（VerifyXXX.go）具体方法。
// 路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
// 路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)
// 路由命名原则：使用3层路由。第1层：接口版本名；第2层：不同拦截器下的不同空间命名；第3层：目标函数handler。
func RegisterRoutes(route *gin.Engine) {


	// ==系统必要路由==
	// 404路由
	route.NoRoute(Middleware.HttpCors(), Middleware.HttpLimiter(2), func (ctx *gin.Context) {
		var url string = ctx.Request.Host + ctx.Request.URL.Path
		var IP string = ctx.ClientIP()
		Kit.Error("404路由：" + url, IP)
		ctx.JSONP(http.StatusNotFound, map[string]interface{}{
			"state": 404,
			"msg": "未定义此名称的路由名",
			"content": gin.H{
				"url": url,
				"ip": IP,
			},
		})
	}, Middleware.HttpAbort)
	// 默认根路由
	route.Any("/",  Middleware.HttpCors(), Middleware.HttpLimiter(2), func (ctx *gin.Context) {
		Kit.Log("进入了默认根路由", ctx.ClientIP())
		ctx.JSONP(http.StatusForbidden, map[string]interface{}{
			"state": 403,
			"msg": "指定路由名后才可访问",
			"content": map[string]interface{}{
				"gl_version": Common.ServerInfo["gl_version"],
				"go_version": Common.ServerInfo["go_version"],
				"timezone": Common.ServerInfo["timezone"],
			},
		})
	}, Middleware.HttpAbort)
	route.Static("/static", Common.ServerInfo["go_path"] + "views/static/") // 多静态文件的主文件夹
	route.StaticFile("/favicon.ico", Common.ServerInfo["go_path"] + "favicon.ico") // 单个静态文件


	// ==接口分组==
	// 访问：域名/gen1/xxx/xxx
	gen1 := route.Group("/gen1/", Middleware.HttpCors())
	{
		//
		user := gen1.Group("/user/", Middleware.HttpLimiter(2), Gen1.VerifyGen1User)
		{
			user.Any("list_user", Gen1User.ListUser)
			user.Any("that_user", Gen1User.ThatUser)
			user.Any("add_user", Gen1User.AddUser)
			user.Any("update_user", Gen1User.UpdateUser)
			user.Any("del_user", Gen1User.DelUser)
			user.Any("clear_user", Gen1User.ClearUser)
		}

		//
		app := gen1.Group("/app/", Middleware.HttpLimiter(2), Gen1.VerifyGen1App)
		{
			app.Any("test1", Gen1App.Test1)                      // 空路由
			app.Any("tpl", Gen1App.Tpl)                          // 模版输出
			app.Any("tpl_index", Gen1App.TplIndex)               // 模版输出
			app.Any("api", Gen1App.Api)                          // 接口输出-简单数据
			app.Any("api2", Gen1App.Api2)                        // 直接接口输出-复杂数据
			app.Any("init", Gen1App.Test2Run)                    // 直接接口输出-复杂数据
			app.Any("upload_form_file", Gen3Open.UploadFormFile) // 直接接口输出-复杂数据
			app.Any("redis_set", Gen1App.RedisSet)
			app.Any("redis_list", Gen1App.RedisList)
		}

	}


	// ==接口分组==
	// 访问：域名/gen3/xxx/xxx
	gen3 := route.Group("/gen3/", Middleware.HttpCors())
	{
		//
		gen3.Any("app/get_app_token", Middleware.HttpLimiter(2), Gen3.VerifyOpen, Gen3App.GetAppToken, Middleware.HttpAbort)
		gen3.Any("user/user_login", Middleware.HttpLimiter(2), Gen3.VerifyOpen, Gen3User.UserLogin, Middleware.HttpAbort)

		//
		open := gen3.Group("/open/", Middleware.HttpLimiter(2), Gen3.VerifyOpen)
		{
			open.Any("upload_form_file", Gen3Open.UploadFormFile, Middleware.HttpAbort)
		}

		//
		app := gen3.Group("/app/", Middleware.HttpLimiter(4), Gen3.VerifyApp)
		{
			app.Any("list_user", Gen3App.ListUser, Middleware.HttpAbort)
			app.Any("that_user", Gen3App.ThatUser, Middleware.HttpAbort)
			app.Any("that_g_user", Gen3App.ThatGUser, Middleware.HttpAbort)
		}

		//
		user := gen3.Group("/user/", Middleware.HttpLimiter(4), Gen3.VerifyUser)
		{
			user.Any("list_user", Gen3User.ListUser, Middleware.HttpAbort)
			user.Any("that_user", Gen3User.ThatUser, Middleware.HttpAbort)
			user.Any("add_user", Gen3User.AddUser, Middleware.HttpAbort)
			user.Any("update_user", Gen3User.UpdateUser, Middleware.HttpAbort)
			user.Any("del_user", Gen3User.DelUser, Middleware.HttpAbort)
			user.Any("clear_user", Gen3User.ClearUser, Middleware.HttpAbort)
		}

	}


	//
}
