package main

import (
	"fmt"
	"time"
)

// time 时间日期的基本运用
func timeBasicOper() {
	now := time.Now()

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("当前时间：%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// time 中定时器的运用
func testTicker() {
	ticker := time.Tick(10 * time.Second)
	num := 0
	for i := range ticker {
		num++
		fmt.Printf("定时器循环：%v\n", i)
		if num > 5 {
			return
		}
	}
}

// time 时间戳
func testTimeTamp(tamp int64) {
	timeObj := time.Unix(tamp, 0)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("当前时间戳格式化后的 :%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间格式化输出
func timeFormat() {
	now := time.Now()
	// 这里 需要固定的时间格式：golang诞生的时间： 2006-01-02 15：04：00 不然其他的时间是不准的
	fmt.Println("正确格式化后的时间： ", now.Format("2006-1-02 15:04"))
	fmt.Println("错误格式化后的时间： ", now.Format("2020-01-02 15:04:00"))
}

func main() {
	timeBasicOper()
	testTicker()
	timeFormat()
	testTimeTamp(time.Now().Unix())
}
