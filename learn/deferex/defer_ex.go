package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) //defer是后进先出，类似堆栈
	}
}
