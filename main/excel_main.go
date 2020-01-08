package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"goProject/custtype"
	"goProject/util"
	"strconv"
	"time"
)

type FinanceDetail struct {
	OrderNo    string        `json:"order_no"`
	NickName   string        `json:"nick_name"`
	Money      float64       `json:"money"`
	IncomeType string        `json:"income_type"`
	Income     float64       `json:"income"`
	TradeTime  custtype.Time `json:"trade_time"`
	PayType    string        `json:"pay_type"`
}

func main() {
	result := make([]FinanceDetail, 0)
	//导入到excel
	columns := make([]util.TableColumn, 0)
	columns = append(columns, util.TableColumn{Name: "订单号", Width: 10, Site: "A1"})
	columns = append(columns, util.TableColumn{Name: "消费者", Width: 30, Site: "B1"})
	columns = append(columns, util.TableColumn{Name: "消费金额", Width: 20, Site: "C1"})
	columns = append(columns, util.TableColumn{Name: "收入类型", Width: 20, Site: "D1"})
	columns = append(columns, util.TableColumn{Name: "收入金额", Width: 20, Site: "E1"})
	columns = append(columns, util.TableColumn{Name: "交易时间", Width: 20, Site: "F1"})
	columns = append(columns, util.TableColumn{Name: "支付方式", Width: 20, Site: "G1"})
	//f :=
	util.CreateNewExcel(&columns, len(result)+1, func(excel *excelize.File) {
		for i, finance := range result {
			iStr := strconv.Itoa(i + 2)
			excel.SetCellValue("Sheet1", "A"+iStr, finance.OrderNo)
			excel.SetCellValue("Sheet1", "B"+iStr, finance.NickName)
			excel.SetCellValue("Sheet1", "C"+iStr, finance.Money)
			excel.SetCellValue("Sheet1", "D"+iStr, finance.IncomeType)
			excel.SetCellValue("Sheet1", "E"+iStr, finance.Income)
			excel.SetCellValue("Sheet1", "F"+iStr, time.Time(finance.TradeTime).Format("2006-01-02 15:04:05"))
			excel.SetCellValue("Sheet1", "G"+iStr, finance.PayType)
		}
	})
	////f.SaveAs(fileName)
	//fileName := url.QueryEscape(fmt.Sprintf("收入明细%v.xlsx", time.Now().Unix()))
	//c.Writer.Header().Add("content-type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	//err = f.Write(c.Writer)
}
