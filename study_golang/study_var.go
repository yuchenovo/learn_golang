package main

import (
	"errors"
	"fmt"
)

type Weekday int

// 类似于枚举
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

/*
交换
*/
func change(a int, b int) {
	a, b = b, a
	fmt.Println(a, b)
}

/*
返回匿名变量
匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
*/
func anymos() (int, int) {
	return 300, 400
}
func add(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("number error")
	}
	d := a / b
	return d, nil
}
func add1(a int, b int) int {
	//defer 要在方法内全部代码执行完毕后才可执行，后注册的defer先执行
	//defer后面是语句时，在注册时候就计算出来 如果是匿名函数，那么在函数计算时才执行
	d := a + b
	defer fmt.Println("111", d)
	fmt.Println(d)
	defer fmt.Println("222", d)
	defer func() {
		fmt.Println("333", d)
	}()
	d = 100
	return d
}
func main() {
	var (
		a = 100
		b = 200
		c int
		d int
		e float32 = .71828
		f         = .71828002     //优先64
		g         = complex(3, 4) //复数
	)
	change(a, b)
	c, _ = anymos()
	_, d = anymos()
	fmt.Println(c, d)
	fmt.Printf("%f\n", e)
	//强制类型转换 只有相同底层类型的变量之间可以进行相互转换（如将 int16 类型转换成 int32 类型）
	//不同底层类型的变量相互转换时会引发编译错误（如将 bool 类型转换为 int 类型）
	fmt.Println(int(e))
	fmt.Printf("%.3f\n", f)
	fmt.Println(g)
	fmt.Println(real(g))
	fmt.Println(imag(g))
	fmt.Println(Saturday)
	fmt.Println(add(1, 0))
	fmt.Println(add1(6, 7))
}
