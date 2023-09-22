package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	//learnArr()
	//learnSlice()
	//learnMap()
	learnList()
}
func learnList() {
	l := list.New()

	// 尾部添加
	l.PushBack("canon")

	// 头部添加
	l.PushFront(67)

	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")

	// 在fist之后添加high
	l.InsertAfter("high", element)

	// 在fist之前添加noon
	l.InsertBefore("noon", element)

	// 使用
	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
func learnMap() {
	//key		value
	var m map[string]int
	//m = make(map[string]int, 100)
	m = map[string]int{"a": 3, "d": 4, "b": 5}
	m["c"] = 9
	fmt.Println(m)
	delete(m, "c")
	if v, exists := m["c"]; exists {
		fmt.Println(v)
	} else {
		fmt.Println("not exist char c")
	}
	var sceneList []string
	for k := range m {
		sceneList = append(sceneList, k)
	}
	fmt.Println(sceneList)
	sort.Strings(sceneList)
	fmt.Println(sceneList)
}
func learnSlice() {
	arr := make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 1, 2, 3
	brr := arr
	brr[0] = 4
	//fmt.Println(arr[0])
	//crr := append(arr, 8)
	//fmt.Println(arr, crr)
	fmt.Print("old:", arr, "\n")
	//arr = append(arr[:0], arr[1:]...) // 删除开头1个元素
	//arr = append(arr[:0], arr[N:]...) // 删除开头N个元素
	arr = arr[:copy(arr, arr[1:])] // 删除开头1个元素
	//arr = arr[:copy(arr, arr[N:])] // 删除开头N个元素
	fmt.Print("new:", arr, "\n")

	drr := make([]int, 2, 5)
	//切片复制
	copy(drr, arr[0:2])
	fmt.Println(drr)
	var highRiseBuilding [30]int

	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}

	// 区间
	fmt.Println(highRiseBuilding[10:15])

	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])

	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])
}
func learnArr() {
	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素

	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	//初始化数组
	//var q [3]int = [3]int{1, 2, 3}
	var r = [3]int{1, 2}
	fmt.Println(r[2]) // "0"
	//在数组的定义中，如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算
	//因此，上面数组 q 的定义可以简化为：
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"
	//数组的长度是数组类型的一个组成部分，因此 [3]int 和 [4]int 是两种不同的数组类型，数组的长度必须是常量表达式
	//因为数组的长度需要在编译阶段确定。
	//q := [3]int{1, 2, 3}
	// 编译错误：无法将 [4]int 赋给 [3]int
	//q = [4]int{1, 2, 3, 4}
}
