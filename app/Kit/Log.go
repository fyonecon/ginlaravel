package Kit

import (
	"ginlaravel/app/Common"
	"os"
	"time"
)

// 记录日志
func Log(_txt string, _ip string) {

	// 文件信息
	fileName := Common.GetTimeDate("Y-m-d") + ".log"
	filePath := Common.ServerInfo["storage_path"] + "log/"
	file, err := os.OpenFile(filePath+fileName, os.O_CREATE | os.O_APPEND |os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	// 延迟关闭文件
	defer file.Close()
	//
	date := Common.GetTimeDate("Y-m-d H:i:s")
	txt := "【" + date + "】" + _ip + "：" + _txt
	// 写入文件内容
	file.WriteString(txt + "\n")

	// 删除老文件，默认存放2个月
	_oldMonth := -2
	day21 := time.Now().AddDate(0, _oldMonth, -1).Format("2006-01-02")
	day22 := time.Now().AddDate(0, _oldMonth, -2).Format("2006-01-02")
	day23 := time.Now().AddDate(0, _oldMonth, -3).Format("2006-01-02")
	day24 := time.Now().AddDate(0, _oldMonth, -4).Format("2006-01-02")
	day25 := time.Now().AddDate(0, _oldMonth, -5).Format("2006-01-02")
	day26 := time.Now().AddDate(0, _oldMonth, -6).Format("2006-01-02")
	day27 := time.Now().AddDate(0, _oldMonth, -7).Format("2006-01-02")
	dayArray := [7]string{day25, day24, day23, day22, day21, day26, day27}
	for d:=0; d<len(dayArray); d++ {
		err := os.Remove(filePath + dayArray[d]+".log")
		// os.Remove(filePath +"2021-04-10.log")
		// fmt.Println(filePath + dayArray[d]+".log")
		if err != nil {
			// 删除失败
		}
	}

}
