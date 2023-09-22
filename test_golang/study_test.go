package test_golang

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

type Student1 struct {
	Name string
	Age  int
	Sex  bool
}
type Class1 struct {
	Id       string
	Students []Student1
}

var (
	s = Student1{"xyh", 23, true}
	c = Class1{"1班", []Student1{s, s, s}}
)

func TestJson(t *testing.T) {
	bytes, err := sonic.Marshal(c)
	if err != nil {
		t.Fail()
	}
	var c1 Class1
	err1 := sonic.Unmarshal(bytes, &c1)
	if err1 != nil {
		t.Fail()
	}
	fmt.Println("success")
}
func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := sonic.Marshal(c)
		var c1 Class1
		sonic.Unmarshal(bytes, &c1)
	}
}
func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(c)
		var c1 Class1
		json.Unmarshal(bytes, &c1)
	}
}
