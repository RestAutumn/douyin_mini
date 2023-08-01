package util1

import (
	"log"
	"time"
)

// 北京时区的时间戳转时间
func UnixToTime(unixTime int64) time.Time {
	timeObj := time.Unix(unixTime, 0)

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("time loadLocation setted failed")
	}

	timeObj = timeObj.In(location)

	return timeObj
}

func TimeNow() time.Time {
	return UnixToTime(time.Now().Unix())
}

// 北京时区的时间转时间戳
func TimeToUnix(time0 time.Time) int64 {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("time loadLocation setted failed")
	}
	time0 = time0.In(location)

	timestamp := time0.Unix()

	return timestamp
}

// 2023-07-31 12:00:18 格式的时间转时间戳
func TimeStringToUnix(time0 string) int64 {
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, time0)
	if err != nil {
		log.Fatal("Error parsing time:", err)
	}

	return t.Unix()
}

func TimeFormat(time0 time.Time) string {
	return time0.Format("2006-01-02 15:04:05")
}
