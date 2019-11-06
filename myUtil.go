package main

import (
	"time"
	"errors"
	"fmt"
)

//获取两个日期间的所有日期
func GetDayListByDateDiff(startTime string, endTime string) ([]string, error) {
	dates := make([]string, 0)
	startDate, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return dates, errors.New("日期参数有误")
	}
	endDate, err := time.Parse("2006-01-02", endTime)
	if err != nil {
		return dates, errors.New("日期参数有误")
	}
	day := TimeSub(startDate, endDate)
	for i := 0; i < day; i++ { //循坏天数 在开始时间基础上 计算这两个日期之间的具体日期 包含开始日期 不包含结束日期
		if i == 0 {
			dates = append(dates, startTime)
		} else {
			dates = append(dates, startDate.AddDate(0, 0, i).Format("2006-01-02")) //加天数
		}
	}
	return dates, nil
}

//两个日期 相差天数
func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t2.Sub(t1).Hours() / 24)
}

func main() {
	arr,_:=GetDayListByDateDiff("2019-10-30","2019-11-06")
	fmt.Println(arr)
}
