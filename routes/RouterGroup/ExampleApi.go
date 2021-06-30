package RouterGroup

/*
路由访问原则：宽进严出，所以都用Any，在拦截器里面拦截（VerifyXXX.go）具体请求事件。
路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)
路由命名原则：推荐使用4层路由。第1层：api类还是web类；第2层：接口版本名；第3层：不同拦截器下的不同空间命名；第4层：目标函数handler。
*/

import (
	"ginvel.com/app/Http/Controllers/Example"
	"ginvel.com/app/Http/Controllers/Example/Captcha"
	"ginvel.com/app/Http/Controllers/Example/ControllerGorm"
	"ginvel.com/app/Http/Controllers/Example/ControllerMySQL"
	"ginvel.com/app/Http/Controllers/Example/Recommend"
	"ginvel.com/app/Http/Controllers/Example/Redis"
	"ginvel.com/app/Http/Controllers/Example/Segment"
	"ginvel.com/app/Http/Controllers/Example/Test"
	"ginvel.com/app/Http/Controllers/Example/WebSocket"
	"ginvel.com/app/Http/Middlewares"
	"github.com/gin-gonic/gin"
)

// ExampleApi 面向模版。访问：你的域名/api/空间命名/具体方法
func ExampleApi(route *gin.Engine)  {
	// api分组路由
	api := route.Group("/api/", Middlewares.HttpCorsApi)
	{
		//
		example := api.Group("/example/")
		{
			//
			mysql := example.Group("/mysql/", Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(90), Example.VerifyExample)
			{
				mysql.Any("list_user", ControllerMySQL.ListUser)
				mysql.Any("that_user", ControllerMySQL.ThatUser)
				mysql.Any("add_user", ControllerMySQL.AddUser)
				mysql.Any("update_user", ControllerMySQL.UpdateUser)
				mysql.Any("del_user", ControllerMySQL.DelUser)
				mysql.Any("clear_user", ControllerMySQL.ClearUser)
			}

			//
			gorm := example.Group("/gorm/", Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(90), Example.VerifyExample)
			{
				gorm.Any("list_user", ControllerGorm.ListUser)
				gorm.Any("that_user", ControllerGorm.ThatUser)
				gorm.Any("add_user", ControllerGorm.AddUser)
				gorm.Any("update_user", ControllerGorm.UpdateUser)
				gorm.Any("del_user", ControllerGorm.DelUser)
				gorm.Any("clear_user", ControllerGorm.ClearUser)
			}

			//
			redis := example.Group("/redis/", Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(90), Example.VerifyExample)
			{
				redis.Any("redis_set", Redis.RedisSet)
				redis.Any("redis_list", Redis.RedisList)
			}

			//
			test := example.Group("/test/", Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(90), Example.VerifyExample)
			{
				test.Any("test1", Test.Test1)
				test.Any("test2", Test.Test2)
			}

			// 图形验证码
			cap := example.Group("/cap/", Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(90), Example.VerifyExample)
			{
				cap.Any("cap", Captcha.Captcha)
			}

			//
			seg := example.Group("/seg/", Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(90), Example.VerifyExample)
			{
				seg.Any("save_seg1", Segment.SaveSeg1)
				seg.Any("search_seg1", Segment.SearchSeg1)
				seg.Any("seg2", Segment.Seg2)
			}

			//
			segment := example.Group("/rec/", Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(90), Example.VerifyExample)
			{
				segment.Any("save_seg1", Recommend.SaveSeg1)
				segment.Any("search_seg1", Recommend.SearchSeg1)
			}

			//
			socket := example.Group("/socket/", Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(90), Example.VerifyExample)
			{
				socket.Any("ping1", WebSocket.Ping1)
			}

		}

		// ====
	}
}
