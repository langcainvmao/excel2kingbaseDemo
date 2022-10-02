package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	_ "kingbase/gokb"
)

//变量

func main() {

	// 读取excel 的内容到数组
	ReadExcel("123.xlsx")

	// 将数组写入到数据库

}

func ReadExcel(excelName string) {
	excelContent := make([][]string, 1000) // 不适用make 会报读

	//for i := 0 ; i< 20; i++ {
	//	excelContent[i] = make([]string,20)
	//}

	// 打开 excel 文件
	f, err := excelize.OpenFile(excelName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭文件
	defer func() {
		// 关闭工作簿
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取表 sheet1 上所有单元格的值
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取每一行的信息
	var rowNum int
	var row []string
	for rowNum, row = range rows {
		if row == nil {
			continue
		}
		for _, colContent := range row {
			if colContent == "" {
				colContent = "null"
			}
			// 要使用append给二维数组中的一维添加， 如果使用excelContent[rowNum][colNum] = colContent 复制会报错。
			excelContent[rowNum] = append(excelContent[rowNum], colContent)
		}
	}
	for i := 0; i <= rowNum; i++ {
		for _, content := range excelContent[i] {
			fmt.Printf("%#v\t", content)
		}
		fmt.Println()
	}
}
