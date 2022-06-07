package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("go")
			ch <- i
		}
		//close(ch) need close channel
	}()

	for n := range ch {
		fmt.Println("main", n)
		//println(n) // 0 1 2 3 4 5 6 7 8 9 deadlock
	}
}