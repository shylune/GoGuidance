package main

import (
	"fmt"
	"os"
)

func main() {
	who := "world"
	if len(os.Args) > 1 {
		who = os.Args[1]
	}

	fmt.Println("Hello", who)
}

// golang 初体验
