package main

import (
	"errors"
	"fmt"
)

func main() {
	//simplefuncTest()
	//multiValFuncTest()
	//passValPeraTest()
	//passPeraPtrTest()
}

func funcTypeTest(){
	fmt.Println(type(simplefuncTest))
}

func simplefuncTest() {
	b, l := max(2, 9)
	println(b, l)
	h, j := SumAndProduct(b, l)
	println(h, j)
}

func multiValFuncTest() {
	m, err := varaPeraMax(4, 6, 8, 34)
	if err == nil {
		println("max val:", m)
	} else {
		fmt.Println(err)
	}
}

//函数参数值传递
func passValPeraTest() {
	x := 3
	fmt.Println("x = ", x)    // 应该输出 "x = 3"
	x1 := add(x)              //调用add(x)
	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)    // x的值未改变，输出"x = 3"
}

//函数参数指针（引用）传递
//channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针
func passPeraPtrTest() {
	x := 3
	fmt.Println("x = ", x)    // 应该输出 "x = 3"
	x1 := addPtr(&x)          //调用addPtr(&x)，传递x变量的指针（内存地址）
	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)    // x的值已改变，输出"x = 4"
}

func max(a, b int) (int, int) {
	if a >= b {
		return a, b
	}
	return b, a
}

func varaPeraMax(args ...int) (maxval int, err error) {
	if len(args) < 3 {
		err = errors.New("length of args must greater than 2")
		return
	}
	for _, v := range args {
		maxval, _ = max(maxval, v)
	}
	return
}

//SumAndProduct 计算输入数据的和 和 积
func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

//简单的一个函数，实现了参数+1的操作
func add(a int) int {
	a = a + 1 // 我们改变了a的值
	return a  //返回一个新值
}

//简单的一个函数，实现了参数+1的操作
func addPtr(a *int) int { // 请注意，
	*a = *a + 1 // 修改了a的值
	return *a   // 返回新值
}
