package models

import "time"

// models中封装一些公共的方法，供controller调用
func UnixToDate(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2022-01-22")
}
