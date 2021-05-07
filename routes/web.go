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
	"ginlaravel/app/Http/Middleware"
	"github.com/gin-gonic/gin"
)

// Web 面向模版。访问：你的域名/web/gen代/空间命名/具体方法
func Web(route *gin.Engine) { //
	// web静态文件
	route.Static("/static", Common.ServerInfo["go_path"] + "views/static/") // 多静态文件的主文件夹
	// web分组路由
	web := route.Group("/web/", Middleware.HttpCorsWeb)
	{
		// ====

		//
		gen1 := web.Group("/gen1/")
		{

			//
			app := gen1.Group("/app/", Middleware.HttpLimiter(3), Gen1.VerifyGen1App)
			{
				app.Any("tpl", Gen1App.Tpl)                          // 模版输出
				app.Any("tpl_index", Gen1App.TplIndex)               // 模版输出
			}

		}


		// ====
	}
}
