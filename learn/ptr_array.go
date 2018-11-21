package main

import "fmt"

/*
MAX COUNT
*/
const MAX int = 3

func main() {
	a := [MAX]int{10, 100, 200} //数组
	//a.append(50) //数组没有append方法
	//printSlice(a)
	//a = append(a, 60, 70) //append第一个参数必须是slice
	printSlice(a)

	var ptr [MAX]*int //指向数组的指针
	fmt.Printf("%T\n", ptr)

	var i int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func printSlice(x [MAX]int) {
	fmt.Printf("len=%d cap=%d type=%T\n", len(x), cap(x), x)
}
