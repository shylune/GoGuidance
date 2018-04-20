/*
常量的定义与变量类似， 使用 const 关键字。
常量可以是字符、字符串、布尔或数字类型的值。
常量不能使用 := 语法定义
数值常量是高精度的 _值_ , 一个未指定类型的常量由上下文来决定其类型
*/

package main

import (
	"fmt"
)

const (
	Pi    = 3.14
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const (
		world = "世界"
		truth = true
	)

	fmt.Println("Hello", world)
	fmt.Println("你好", Pi)
	fmt.Println("Go rules?", truth)

	fmt.Println(needInt(Small))
	//	fmt.Println(needInt(Big))	// overflow int
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
