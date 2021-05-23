package Kit
// 生成和读取Excel文件
// github.com/360EntSecGroup-Skylar/excelize/v2

import (
	"fmt"
	"ginvel.com/app/Common"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// MakeExcel 按行生成Excel，仅支持sheet1按行添加数据
// github.com/360EntSecGroup-Skylar/excelize/v2
// 最大支持40列
/*
rowData := [][]interface{}{
		{"name", "phone", "age"},
		{"张三1", "1231", 23},
		{"张三2", "1232", 24},
		{"张三3", "1233", 25},
		{"合计人数", 3},
	}
excelName := "demo.xlsx"
*/
func MakeExcel(rowData [][]interface{}, excelName string, filepath string) string {
	if len(filepath) == 0 {
		filepath = Common.ServerInfo["storage_path"] + "cache_file/"
	}

	// 初始参数
	f := excelize.NewFile()

	// 处理数据
	li := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "J", "K",
		"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE",
		"AF", "AG", "AH", "AJ", "AK", "AL", "AM", "AN", "AO", "AP",
	}
	fmt.Println(li)
	for n:=0; n<len(rowData); n++ { // 每列
		theRow := rowData[n]
		for l:=0; l<len(theRow); l++ { // 每行
			if l < len(li) {
				theLi := li[l] + Common.IntToString(int64(n+1)) // 构建每行坐标=列字母+行标号，默认从第1行开始
				theValue := theRow[l]
				f.SetCellValue("Sheet1", theLi, theValue)
			}else {
				fmt.Println("Excel列数超范围")
				break
			}
		}
	}

	// 生成Excel文件
	// Set active sheet of the workbook.
	f.SetActiveSheet(1)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(filepath + excelName); err != nil {
		fmt.Println(err)
		return ""
	}else {
		return excelName
	}
}

// ReadExcel 按行读取Excel
func ReadExcel(excelName string, filepath string) (rows [][]string, err error) {
	if len(filepath) == 0 {
		filepath = Common.ServerInfo["storage_path"] + "cache_file/"
	}

	// 解析文件
	f, err := excelize.OpenFile(filepath + excelName)
	if err != nil {
		fmt.Println(err)
		return rows, err
	}

	// 按行获取，默认取sheet1表
	_rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return rows, err
	}

	return _rows, err // 返回行
}