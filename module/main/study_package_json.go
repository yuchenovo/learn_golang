package main

import (
	"fmt"
	"github.com/bytedance/sonic"
)

type Student struct {
	Name string
	Age  int
	Sex  bool
}
type Class struct {
	Id       string
	Students []Student
}

func main() {
	s := Student{"xyh", 23, true}
	c := Class{"1班", []Student{s, s, s}}
	//json序列化
	//bytes, err := json.Marshal(c)
	//sonic序列化
	bytes, err := sonic.Marshal(c)
	if err != nil {
		fmt.Println("Serialization failed:", err)
		return
	} else {
		fmt.Println(string(bytes))
	}
	var c1 Class
	////json反序列化：想要修改结构体内部，必须要传递指针，否则只在函数内部生效
	err1 := sonic.Unmarshal(bytes, &c1)
	if err1 != nil {
		fmt.Println("Deserialization failed:", err1)
		return
	}
	fmt.Printf("%+v\n", c1)
}
