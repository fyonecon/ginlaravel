package TaskInterval

import (
	"log"
)

// TaskInterval 全局定时器，默认精度20s/次
func TaskInterval(intervalId int, num int, timeout string) {
	log.Println("全局定时器已开启>>>", " 定时器ID=", intervalId, " 运行次数num=", num, " 时间区间=", timeout)


}
