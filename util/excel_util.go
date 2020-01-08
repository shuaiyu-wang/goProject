package util

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

type TableColumn struct {
	Name  string  //列名称
	Width float64 //表格宽度
	Site  string  //列位置
}

//创建excel导出文件
func CreateNewExcel(columnList *[]TableColumn, rowCount int, writerContent func(excel *excelize.File)) *excelize.File {
	f := excelize.NewFile()
	endColumn := string(rune(len(*columnList) + 64))
	//设置表头
	for _, item := range *columnList {
		f.SetCellValue("Sheet1", item.Site, item.Name)
		column := item.Site[0:1]
		f.SetColWidth("Sheet1", column, column, item.Width)
	}
	f.SetRowHeight("Sheet1", 1, 30) //设置表头高度
	style, _ := f.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"color":"#ffffff"},
    "fill":{"type":"pattern","color":["#0083FE"],"pattern":1}}`)
	endColumnSite := endColumn + "1"
	f.SetCellStyle("Sheet1", "A1", endColumnSite, style) //设置表头样式
	//设置内容对齐样式
	contentStyle, _ := f.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	vCell := endColumn + strconv.Itoa(rowCount)
	f.SetCellStyle("Sheet1", "A2", vCell, contentStyle)
	//设置内容行高
	for i := 2; i <= rowCount; i++ {
		f.SetRowHeight("Sheet1", i, 30)
	}
	//写入内容
	writerContent(f)
	return f
}
