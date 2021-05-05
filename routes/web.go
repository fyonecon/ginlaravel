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

// Web 路由访问原则：宽进严出，所以都用Any，在拦截器里面拦截（VerifyXXX.go）具体请求事件。
// 路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
// 路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)
// 路由命名原则：推挤使用4层路由。第1层：api类还是web类；第2层：接口版本名；第3层：不同拦截器下的不同空间命名；第4层：目标函数handler。
func Web(route *gin.Engine) { // 面向模版。访问：域名/api/genx/xxx/xxx
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
