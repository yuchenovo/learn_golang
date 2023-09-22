package main

import (
	"fmt"
	"time"
)

// 时间格式化固定格式
const (
	DATE = "2006-01-02"
	TIME = "2006-01-02 15:04:05"
)

func main() {
	t0 := time.Now()
	fmt.Println(t0.Unix())
	time.Sleep(50 * time.Millisecond)
	t1 := time.Now()
	diff := t1.Sub(t0)
	fmt.Println(diff.Milliseconds())
	fmt.Println(time.Since(t0).Milliseconds())
	d := time.Duration(2 * time.Second)
	t2 := t0.Add(d)
	fmt.Println(t2.Unix())
	fmt.Println(t0.Format(DATE))
	s := t0.Format(TIME)
	fmt.Println(s)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t3, _ := time.ParseInLocation(TIME, s, loc)
	fmt.Println(t3.Unix())
}
