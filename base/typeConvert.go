package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := int(f)
	fmt.Printf("%v(%T)\n", f, f)
	fmt.Printf("%v(%T)\t%v(%T)\t%v(%T)\n", x, x, y, y, z, z)
}

/*
在定义一个变量但不指定其类型时（使用没有类型的 var 或 := 语句）， 变量的类型由右值推导得出
*/
