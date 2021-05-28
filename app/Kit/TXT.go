package Kit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TxtReadLine 按行读txt文件
func TxtReadLine(filepath string, filename string) (txtArray []string) {
	// 计数
	lineNumHave := 0 // 有值
	lineNumNull := 0 // 空值

	// 打开txt文件
	file, err := os.Open(filepath + filename)
	if err != nil{
		fmt.Println("无效的txt文件", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// 按行处理txt
	for scanner.Scan(){
		lineTxt := strings.TrimSpace(scanner.Text())

		if len(lineTxt) == 0 {
			lineTxt = "null"
			lineNumNull ++
		}else {
			lineNumHave ++
		}

		txtArray = append(txtArray, lineTxt)
	}

	return
}
