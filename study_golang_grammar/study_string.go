package main

import "fmt"

func main() {
	//字符串是不可变的byte数组
	s := "你好golang"
	for i, e := range s {
		fmt.Printf("%d  %c\n", i, e)
	}
	fmt.Println(len(s))
	//rune切片把字符串的每个字符转化为1个字节，否则输出的汉字为3字节
	arr := []rune(s)
	fmt.Println(arr)
	for i, e := range arr {
		fmt.Printf("%d  %c\n", i, e)
	}
	fmt.Println(len(arr))
}
