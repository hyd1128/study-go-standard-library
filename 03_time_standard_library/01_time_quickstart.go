package main

import (
	"fmt"
	"time"
)

// time类型
// func timeDemo() {
// 	now := time.Now()
// 	fmt.Printf("current time: %v \n", now)
// }

// 2. 时间戳
func timestampDemo() {
	now := time.Now()            // 获取当前时间
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func timeParse() {
	t, err := time.Parse("2006-01-02 15:04:05", "2022-07-28 18:06:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", now.Format("2006/01/02 15:04:05"), loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
}

func timeFormat() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分05秒
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("2006"))
}

// unix时间戳
func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     // 年
	month := timeObj.Month()   // 月
	day := timeObj.Day()       // 日
	hour := timeObj.Hour()     // 小时
	minute := timeObj.Minute() // 分钟
	second := timeObj.Second() // 秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func main() {
	// 显示当前时间
	// timeDemo()
	// now := time.Now()
	// year := now.Year()
	// month := now.Month()
	// day := now.Day()
	// hour := now.Hour()
	// minute := now.Minute()
	// second := now.Second()
	// fmt.Printf("当前时间: %v-%v-%v %v:%v:%v", year, month, day, hour, minute, second)

	// timestampDemo()

	// timeParse()

	// timeFormat()

	// now := time.Now()
	// timestampDemo2(now.Unix())

	// 1纳秒表示方式
	// fmt.Printf("%v", 1*time.Nanosecond)

	// 时间的计算
	// add 1 hour
	// now := time.Now()
	// later := now.Add(time.Hour)
	// fmt.Println(later)

	// 计算两个时间之间的差值
	// now := time.Now()
	// later := now.Add(time.Hour)
	// ret := later.Sub(now)
	// fmt.Println(ret)

	// 判断时间1和时间2的关系
	timeOne := time.Now()
	timeTwo := time.Now()
	result := timeOne.Equal(timeTwo)
	fmt.Printf("%v %T", result, result)

	timeThree := timeOne.Add(-time.Hour)
	fmt.Printf("%v", timeThree)
}
