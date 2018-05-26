package main

import "fmt"

/*
MAX COUNT
*/
const MAX int = 3

func main() {
	a := [MAX]int{10, 100, 200}
	// a.append(50)
	//printSlice(a)
	// a = append(a, 60, 70)
	// printSlice(a)
	var i int
	var ptr [MAX]*int
	fmt.Printf("%T\n", ptr)

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
