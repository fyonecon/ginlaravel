package driver
// 提示和安装fresh热更服务

import (
	"ginvel.com/app/Common"
	"log"
	"os"
	"os/exec"
)

func InitFresh()  {

	log.Println("检测fresh热更服务...")

	// 文件信息
	fileName := "having-fresh.log"
	filePath := Common.ServerInfo["storage_path"] + "fresh/" + fileName
	freshTips := "（提示：如果没有安装fresh，请先删除'/storage/fresh/'目录下的（所有）文件，然后再次运行「 go run main.go 」）"
	// 判断文件是否存在
	hasFile, err := Common.HasFile(filePath)
	if !hasFile {
		// 检查fresh是否已经安装
		cmd := exec.Command("fresh")
		err := cmd.Start()
		if err != nil {
			log.Println(err)
			log.Println("请运行安装fresh热更服务，请手动运行如下命令：\n go get -u github.com/pilu/fresh \n")
			os.Exit(200)
			//return
		}else {
			// 创建文件
			file, err := os.OpenFile(filePath, os.O_CREATE | os.O_APPEND |os.O_WRONLY, 0666)
			if err != nil {
				log.Println("fresh服务未启动，已跳过", err)
				//panic(err)
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
		log.Println("fresh热更服务已初始化 >>> ", hasFile, err, freshTips)
	}

}
