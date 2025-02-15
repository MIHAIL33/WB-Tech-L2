package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {

	defer fmt.Println("111") // ^
	defer fmt.Println("222") // |

	fmt.Println(test()) // 2
	fmt.Println(anotherTest()) // 1
}