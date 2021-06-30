package RouterGroup

import (
	"ginvel.com/app/Http/Controllers/Gen3"
	"ginvel.com/app/Http/Controllers/Gen3/Gen3App"
	"ginvel.com/app/Http/Controllers/Gen3/Gen3Open"
	"ginvel.com/app/Http/Controllers/Gen3/Gen3User"
	"ginvel.com/app/Http/Middlewares"
	"github.com/gin-gonic/gin"
)

func ApiGen3(route *gin.Engine)  {
	// api分组路由
	api := route.Group("/api/", Middlewares.HttpCorsApi)
	{
		//
		gen3 := api.Group("/gen3/")
		{
			//
			gen3.Any("app/get_app_token", Middlewares.HttpLimiter(2), Gen3.VerifyOpen, Gen3App.GetAppToken)
			gen3.Any("user/user_login", Middlewares.HttpLimiter(2), Gen3.VerifyOpen, Gen3User.UserLogin)

			//
			open := gen3.Group("/open/", Middlewares.HttpLimiter(2), Gen3.VerifyOpen)
			{
				open.Any("upload_form_file", Gen3Open.UploadFormFile)
			}

			//
			app := gen3.Group("/app/", Middlewares.HttpLimiter(4), Gen3.VerifyApp)
			{
				app.Any("list_user", Gen3App.ListUser)
				app.Any("that_user", Gen3App.ThatUser)
				app.Any("that_g_user", Gen3App.ThatGUser)

			}

			//
			user := gen3.Group("/user/", Middlewares.HttpLimiter(4), Gen3.VerifyUser)
			{
				user.Any("list_user", Gen3User.ListUser)
				user.Any("that_user", Gen3User.ThatUser)
				user.Any("add_user", Gen3User.AddUser)
				user.Any("update_user", Gen3User.UpdateUser)
				user.Any("del_user", Gen3User.DelUser)
				user.Any("clear_user", Gen3User.ClearUser)
			}

		}


		// ====
	}
}
