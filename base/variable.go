package main

import "fmt"

var c, python, java bool
var i, j = 1, 2

//t := "test"

func main() {
	var k int
	cpp, py, php := true, false, "no!" // short declaration

	fmt.Println(c, python, java, k)
	fmt.Println(i, j)
	fmt.Println(cpp, py, php)

	//	fmt.Println(t)
}

/*
var 语句定义了一个变量的列表, 跟函数的参数列表一样，类型在后面, 可以定义在包或函数级别
变量 定义可以包含初始值, 如果初始化是使用表达式， 则可以省略类型, 变量从初始值中获得类型

在函数中， `:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义
在函数外, 每个语句都必须以关键字开始（`var`、`func`、等等）， `:=` 结构不能使用在函数外
*/
