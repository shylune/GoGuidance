package main

import (
	"fmt"
)

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add1(5, 9), add2(2, 4))
	a, b := swap("value", "return")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

/*
注意函数形参,形参类型的写法, 类型在变量名 _之后_
当两个或多个连续的函数命名参数是同一类型， 则除了最后一个类型之外，其他都可以省略

函数可以返回任意数量的返回值
返回值可以被命名，并且像变量那样使用, 返回值的名称应当具有一定的意义，可以作为文档使用
没有参数的 return 语句返回结果的__当前值__
*/
