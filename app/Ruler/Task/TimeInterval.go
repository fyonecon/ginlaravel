package Task

import (
	"log"
)

// TimeInterval 全局定时器，默认精度20s/次
func TimeInterval(intervalId int, num int, timeout string) {
	var maxLog int = 5
	if num < maxLog { // 不必全部打印，只打印前几个即可
		log.Println("全局定时器已开启>>>", " 定时器ID=", intervalId, " 运行次数num=", num, " 时间区间=", timeout)
	}else if num == maxLog {
		log.Println("全局定时器日志显示已关闭，定时任务会继续运行。maxLog=", maxLog)
	}

	// 其他定时任务


}
