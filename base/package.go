package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.%s", math.Nextafter(2, 3), "\n")
}

// 程序入口是main包中main()函数，导入包用import，包名与导入路径的最后一个目录一致（按照惯例）
