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
	"ginlaravel/app/Http/Controller/Gen3/Gen3User"
	"ginlaravel/app/Http/Middleware"
	"github.com/gin-gonic/gin"
)

// Api 面向Api。访问：你的域名/web/gen代/空间命名/具体方法
func Api(route *gin.Engine) {
	// api分组路由
	api := route.Group("/api/", Middleware.HttpCorsApi)
	{
		// ====


		gen1 := api.Group("/gen1/")
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
				app.Any("api", Gen1App.Api)                          // 接口输出-简单数据
				app.Any("api2", Gen1App.Api2)                        // 直接接口输出-复杂数据
				app.Any("init", Gen1App.Test2Run)                    // 直接接口输出-复杂数据
				app.Any("upload_form_file", Gen3Open.UploadFormFile) // 直接接口输出-复杂数据
				app.Any("redis_set", Gen1App.RedisSet)
				app.Any("redis_list", Gen1App.RedisList)
			}

		}


		gen3 := api.Group("/gen3/")
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

				app.Any("list_gm_user", Gen3App.ListGMUser, Middleware.HttpAbort)
				app.Any("that_gm_user", Gen3App.ThatGMUser, Middleware.HttpAbort)
				app.Any("add_gm_user", Gen3App.AddGMUser, Middleware.HttpAbort)
				app.Any("update_gm_user", Gen3App.UpdateGMUser, Middleware.HttpAbort)
				app.Any("del_gm_user", Gen3App.DelGMUser, Middleware.HttpAbort)
				app.Any("clear_gm_user", Gen3App.ClearGMUser, Middleware.HttpAbort)
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


		// ====
	}
}
