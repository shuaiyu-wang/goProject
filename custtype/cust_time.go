package custtype

import (
	"time"
)

type Time time.Time

const (
	timeFormart       = "2006-01-02 15:04:05"
	timeFormart_short = "2006-01-02"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if len(string(data)) < 8 {
		return
	}
	var now time.Time
	if len(string(data)) > 14 {
		now, err = time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	} else {
		now, err = time.ParseInLocation(`"`+timeFormart_short+`"`, string(data), time.Local)
	}

	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	rtime := time.Time(t)
	if rtime.Unix() < 0 || rtime.IsZero() {
		return []byte("\"\""), nil
	}
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = rtime.AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

// 可以用于sqlTemplate中,判断时间是否有传值
//func (t Time) IsZero() bool {
//    return time.Time(t).IsZero()
//}
//
//func (t Time) String() string {
//    return time.Time(t).Format(timeFormart)
//}
