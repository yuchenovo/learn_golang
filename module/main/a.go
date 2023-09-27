package main

//_ "g/util" import中 包名前面加下划线 == 不导入此包 但是会执行该包的init函数（init函数可为多个）
import (
	"fmt"
	"g/util"
	"g/util/math"
)

func all() {
	//在不同包下 只有首字母大写 的可被调用
	fmt.Println(util.Name, math.Add(3, 4))
}
