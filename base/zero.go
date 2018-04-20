package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/*
变量在定义时没有明确的初始化时会赋值为_零值_
数值类型为 `0`，
布尔类型为 `false`，
字符串类型为 `""`（空字符串）
*/
