package math

/*
在同一包下函数可互相调用 eg：c.go
在不同包下 只有首字母大写 的可被调用(Add) eg: d.go
*/

func Add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
