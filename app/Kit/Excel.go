package Kit
// 生成和读取Excel文件
// github.com/360EntSecGroup-Skylar/excelize/v2

import (
	"fmt"
	"ginlaravel/app/Common"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// MakeExcel 生成Excel
func MakeExcel() string {
	excelName := Common.MakeSMSCode(8)+".xlsx"

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(Common.ServerInfo["storage_path"] + "cache_file/"+excelName); err != nil {
		fmt.Println(err)
		return ""
	}else {
		return excelName
	}
}

// ReadExcel 读取Excel
func ReadExcel() ([][]string, error) {
	excelName := "Book1.xlsx"

	// 解析文件
	f, err := excelize.OpenFile(Common.ServerInfo["storage_path"] + "cache_file/"+excelName)
	if err != nil {
		fmt.Println(err)
		return [][]string{{}}, err
	}

	// 获取某列
	// Get value from cell by given worksheet name and axis.
	//cell, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//	fmt.Println(err)
	//	return [][]string{{}}, err
	//}
	//fmt.Println(cell)

	// 获取某行
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return [][]string{{}}, err
	}
	//for _, row := range rows {
	//	for _, colCell := range row {
	//		fmt.Print(colCell, "\t")
	//	}
	//}

	return rows, err // 返回行
}