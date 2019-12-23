package test

import (
	"fmt"
	_ "fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx := excelize.NewFile()
	// 创建一个工作表
	index := xlsx.NewSheet("Sheet1")
	// 设置单元格的值
	xlsx.SetCellValue("Sheet1", "A8", "序号")
	xlsx.SetColWidth("Sheet1", "A", "A", 22)
	xlsx.SetCellValue("Sheet1", "B8", "供应商名称")
	xlsx.SetColWidth("Sheet1", "B", "B", 30)
	xlsx.SetCellValue("Sheet1", "C8", "凭证数")
	xlsx.SetColWidth("Sheet1", "C", "C", 15)
	xlsx.SetCellValue("Sheet1", "D8", "金额合计")
	xlsx.SetColWidth("Sheet1", "D", "D", 10)
	xlsx.SetCellValue("Sheet1", "E8", "税金合计")
	xlsx.SetColWidth("Sheet1", "E", "E", 10)
	xlsx.SetCellValue("Sheet1", "F8", "价税合计")
	xlsx.SetColWidth("Sheet1", "F", "F", 10)
	xlsx.SetCellValue("Sheet1", "G8", "一级类目金额统计")
	xlsx.SetColWidth("Sheet1", "G", "G", 10)
	xlsx.SetCellValue("Sheet1", "H8", "")
	xlsx.SetColWidth("Sheet1", "H", "H", 10)
	xlsx.SetCellValue("Sheet1", "I8", "")
	xlsx.SetColWidth("Sheet1", "I", "I", 15)
	xlsx.SetCellValue("Sheet1", "J8", "")
	xlsx.SetColWidth("Sheet1", "J", "J", 15)

	xlsx.SetCellValue("Sheet1", "A9", "")
	xlsx.SetColWidth("Sheet1", "A", "A", 22)
	xlsx.SetCellValue("Sheet1", "B9", "")
	xlsx.SetColWidth("Sheet1", "B", "B", 30)
	xlsx.SetCellValue("Sheet1", "C9", "")
	xlsx.SetColWidth("Sheet1", "C", "C", 15)
	xlsx.SetCellValue("Sheet1", "D9", "")
	xlsx.SetColWidth("Sheet1", "D", "D", 10)
	xlsx.SetCellValue("Sheet1", "E9", "")
	xlsx.SetColWidth("Sheet1", "E", "E", 10)
	xlsx.SetCellValue("Sheet1", "F9", "")
	xlsx.SetColWidth("Sheet1", "F", "F", 10)
	//一级类目金额统计
	xlsx.SetCellValue("Sheet1", "G9", "蔬菜类")
	xlsx.SetColWidth("Sheet1", "G", "G", 10)
	xlsx.SetCellValue("Sheet1", "H9", "米面粉类")
	xlsx.SetColWidth("Sheet1", "H", "H", 10)
	xlsx.SetCellValue("Sheet1", "I9", "调味品类")
	xlsx.SetColWidth("Sheet1", "I", "I", 15)
	xlsx.SetCellValue("Sheet1", "J9", "水果类")
	xlsx.SetColWidth("Sheet1", "J", "J", 15)

	//合并单元格
	xlsx.MergeCell("Sheet1", "G8", "J8")
	xlsx.MergeCell("Sheet1", "A8", "A9")
	xlsx.MergeCell("Sheet1", "B8", "B9")
	xlsx.MergeCell("Sheet1", "C8", "C9")
	xlsx.MergeCell("Sheet1", "D8", "D9")
	xlsx.MergeCell("Sheet1", "E8", "E9")
	xlsx.MergeCell("Sheet1", "F8", "F9")
	//xlsx.MergeCell("Sheet1", "G8", "G9")
	//xlsx.MergeCell("Sheet1", "H8", "H9")

	//测试数据
	xlsx.SetCellValue("Sheet1", "A10", "xxx")
	xlsx.SetCellValue("Sheet1", "B10", "xxx")
	xlsx.SetCellValue("Sheet1", "C10", "xxx")
	xlsx.SetCellValue("Sheet1", "D10", "xxx")
	xlsx.SetCellValue("Sheet1", "E10", "xxx")
	xlsx.SetCellValue("Sheet1", "F10", "xxx")
	xlsx.SetCellValue("Sheet1", "G10", "xxx")
	xlsx.SetCellValue("Sheet1", "H10", "xxx")
	xlsx.SetCellValue("Sheet1", "I10", "xxx")
	xlsx.SetCellValue("Sheet1", "J10", "xxx")

	//设置单元格格式，边框、对齐方式、颜色
	style, _ := xlsx.NewStyle(`{"fill":{"type":"pattern","color":["#7FFFD4"],"pattern":1},"font":{"color":"#000000"},"alignment":{"horizontal":"center","vertical":"center"},
		"border":[{"type":"left","color":"#080808","style":1},{"type":"top","color":"#080808","style":1},
				  {"type":"right","color":"#080808","style":1},{"type":"bottom","color":"#080808","style":1}]}`)
	xlsx.SetCellStyle("sheet1", "A8", "J9", style)

	style, _ = xlsx.NewStyle(`{"fill":{"type":"pattern","color":["#FFFFFF"],"pattern":1},"font":{"color":"#000000"},"alignment":{"horizontal":"center","vertical":"center"},
		"border":[{"type":"left","color":"#080808","style":1},{"type":"top","color":"#080808","style":1},
				  {"type":"right","color":"#080808","style":1},{"type":"bottom","color":"#080808","style":1}]}`)
	xlsx.SetCellStyle("sheet1", "A10", "J10", style)
	// 设置工作簿的默认工作表
	xlsx.SetActiveSheet(index)
	// 根据指定路径保存文件
	err := xlsx.SaveAs("D:\\IOTest\\Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
