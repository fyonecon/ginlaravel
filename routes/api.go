package routes
/*
路由访问原则：宽进严出，所以都用Any，在拦截器里面拦截（VerifyXXX.go）具体请求事件。
路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)
路由命名原则：推荐使用4层路由。第1层：api类还是web类；第2层：接口版本名；第3层：不同拦截器下的不同空间命名；第4层：目标函数handler。
*/

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

import (
	"ginlaravel/app/Http/Controller/Example"
	"ginlaravel/app/Http/Controller/Example/ControllerGorm"
	"ginlaravel/app/Http/Controller/Example/ControllerMySQL"
	"ginlaravel/app/Http/Controller/Example/Redis"
	"ginlaravel/app/Http/Controller/Example/Test"
	"ginlaravel/app/Http/Controller/Gen3"
	"ginlaravel/app/Http/Controller/Gen3/Gen3App"
	"ginlaravel/app/Http/Controller/Gen3/Gen3Open"
	"ginlaravel/app/Http/Controller/Gen3/Gen3User"
	"ginlaravel/app/Middleware"
	"github.com/gin-gonic/gin"
)

// Api 面向Api。访问：你的域名/api/版本/空间命名/具体方法
func Api(route *gin.Engine) {
	// api分组路由
	api := route.Group("/api/", Middleware.HttpCorsApi)
	{
		//
		example := api.Group("/example/")
		{
			//
			mysql := example.Group("/mysql/", Middleware.HttpLimiter(2), Example.VerifyExample)
			{
				mysql.Any("list_user", ControllerMySQL.ListUser)
				mysql.Any("that_user", ControllerMySQL.ThatUser)
				mysql.Any("add_user", ControllerMySQL.AddUser)
				mysql.Any("update_user", ControllerMySQL.UpdateUser)
				mysql.Any("del_user", ControllerMySQL.DelUser)
				mysql.Any("clear_user", ControllerMySQL.ClearUser)
			}

			//
			gorm := example.Group("/gorm/", Middleware.HttpLimiter(2), Example.VerifyExample)
			{
				gorm.Any("list_user", ControllerGorm.ListUser)
				gorm.Any("that_user", ControllerGorm.ThatUser)
				gorm.Any("add_user", ControllerGorm.AddUser)
				gorm.Any("update_user", ControllerGorm.UpdateUser)
				gorm.Any("del_user", ControllerGorm.DelUser)
				gorm.Any("clear_user", ControllerGorm.ClearUser)
			}

			//
			redis := example.Group("/redis/", Middleware.HttpLimiter(2), Example.VerifyExample)
			{
				redis.Any("redis_set", Redis.RedisSet)
				redis.Any("redis_list", Redis.RedisList)
			}

			//
			test := example.Group("/test/", Middleware.HttpLimiter(2), Example.VerifyExample)
			{
				test.Any("test1", Test.Test1)
			}

		}

		//
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
