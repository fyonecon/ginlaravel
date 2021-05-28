package Kit

import (
	"fmt"
	"ginvel.com/app/Common"
	"log"
	"os"
	"time"
)

// Log 记录一般日志，定期自动删除
func Log(_txt string, _ip string) {

	// 创建文件夹
	filepath := Common.ServerInfo["storage_path"]+"log_file/log/"
	dateFile := Common.GetTimeDate("Ymd") + "/"
	saveFilepath := filepath + dateFile
	// 创建日期文件夹
	has, _ := Common.HasFile(saveFilepath)
	if !has {
		err := os.Mkdir(saveFilepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹1=[%v]\n", err)
		}
	}

	// 文件信息
	fileName := "log_" + Common.GetTimeDate("Y-m-d") + ".log"
	filePath := saveFilepath
	file, err := os.OpenFile(filePath+fileName, os.O_CREATE | os.O_APPEND |os.O_WRONLY, 0666)
	if err != nil {
		log.Println("文件创建失败", filePath, err)
		//panic(err)
		return
	}
	// 延迟关闭文件
	defer file.Close()
	//
	date := Common.GetTimeDate("Y-m-d H:i:s")
	txt := "【" + date + "】" + _ip + "：" + _txt
	// 写入文件内容
	_, err = file.WriteString(txt + "\n")
	if err != nil {
		log.Println("文件写入失败", filePath, err)
		return
	}

	// 删除老文件夹
	_oldMonth := -18
	delFile := time.Now().AddDate(0, _oldMonth, -1).Format("20060102")
	delFilepath := filepath + delFile
	err1 := os.RemoveAll(delFilepath)
	if err1 != nil {
		log.Println("老文件夹删除失败=", err1)
	}

}

// Error 记录错误日志
func Error(_txt string, _ip string) {

	// 创建文件夹
	filepath := Common.ServerInfo["storage_path"]+"log_file/error/"
	dateFile := Common.GetTimeDate("Ymd") + "/"
	saveFilepath := filepath + dateFile
	// 创建日期文件夹
	has, _ := Common.HasFile(saveFilepath)
	if !has {
		err := os.Mkdir(saveFilepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹1=[%v]\n", err)
		}
	}

	// 文件信息
	fileName := "error_" + Common.GetTimeDate("Y-m-d") + ".log"
	filePath := saveFilepath
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

	// 删除老文件夹
	_oldMonth := -18
	delFile := time.Now().AddDate(0, _oldMonth, -1).Format("20060102")
	delFilepath := filepath + delFile
	err1 := os.RemoveAll(delFilepath)
	if err1 != nil {
		log.Println("老文件夹删除失败=", err1)
	}

}