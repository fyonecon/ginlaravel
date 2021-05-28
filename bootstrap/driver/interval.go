package driver
// 全局定时器
// 文档：https://blog.csdn.net/weixin_41827162/article/details/117152232

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Ruler/Task"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func InitInterval()  {

	// timeout示例，写法v1的相同：
	// 每12s运行一次："@every 12s" 或 "*/12 * * * * *"
	// 每分钟的第0s执行一次："0 */1 * * * *"
	// 每分钟的第5s、25s、45s各执行一次：5,25,45 * * * * *
	// 每12s执行一次：*/12 * * * * *
	// 每隔1分钟的第0秒执行一次：0 */1 * * * *
	// 每天23:30:00执行一次：0 30 23 * * *
	// 每天凌晨1:00:00执行一次：0 0 1 * * *
	// 每月1号早上6:00:00执行一次：0 0 6 1 * *
	// 在每小时的26分、29分、33分各执行一次：0 26,29,33 * * *
	// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * *
	// 每天十点到十二点每五秒执行一次：*/5 * 10-12 * * *

	var timeout string = "0,30 * * * * *" // 定时器时间区间，默认精度为30s/次：0,30 * * * * *
	var intervalId int // 定时器id

	go func() {
		num := 0 // 运行次数
		// 设置时区
		local, _ := time.LoadLocation(Common.ServerInfo["timezone"])
		interval := cron.New(cron.WithLocation(local), cron.WithSeconds()) // 设置时区并且精度按秒。
		_timeout := timeout
		_intervalId, err := interval.AddFunc(_timeout, func() {
			num++
			//log.Println("全局定时器已开启 >>> ", " 定时器ID=", intervalId, " 运行次数num=", num, " 时间区间=", timeout)
			// 下面调用其他函数
			Task.TimeInterval(intervalId, num, _timeout)

		})
		if err != nil{
			log.Println("全局定时器报错：", " error=", err, " num=", num)
			os.Exit(200)
		}
		intervalId = int(_intervalId)
		log.Println("全局定时器已开启 >>>", " 定时器Id=", intervalId, " 默认精度=", _timeout)
		interval.Start()

		//关闭着计划任务, 但是不能关闭已经在执行中的任务.
		//defer interval.Stop()
		//select{} // 阻塞主线程而不退出

	}()

}
