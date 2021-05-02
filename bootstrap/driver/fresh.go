package driver

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

// 提示和安装fresh热更服务

import (
	"ginlaravel/app/Common"
	"log"
	"os"
	"os/exec"
)

func init()  {

	log.Println("检测fresh热更服务...")

	// 文件信息
	fileName := "having-fresh.log"
	filePath := Common.ServerInfo["storage_path"] + "fresh/" + fileName
	freshTips := "如果没有安装fresh，请先删除'/storage/fresh/'目录下的（所有）文件，然后再次运行「 go run server.go 」"
	// 判断文件是否存在
	hasFile, err := Common.HasFile(filePath)
	if !hasFile {
		// 检查fresh是否已经安装
		cmd := exec.Command("fresh")
		err := cmd.Start()
		if err != nil {
			log.Println(err)
			log.Println("请运行安装fresh热更服务，请手动运行如下命令：\n go get github.com/pilu/fresh \n")
			os.Exit(200)
		}else {
			// 创建文件
			file, err := os.OpenFile(filePath, os.O_CREATE | os.O_APPEND |os.O_WRONLY, 0666)
			if err != nil {
				panic(err)
			}else {
				defer file.Close()
				date := Common.GetTimeDate("Y-m-d H:i:s")
				txt := date
				file.WriteString(txt + "\n")
				file.WriteString(freshTips + "\n")
			}

			log.Println("fresh热更服务已启动 >>> ")
		}

	}else {
		log.Println(hasFile)
		log.Println(err)
		log.Println(freshTips)
	}

}
