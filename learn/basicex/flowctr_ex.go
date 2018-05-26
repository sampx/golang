package main

import "fmt"

func main() {
	//ifTest()
	//gotoTest()
	//forTest()
	//whileTest()
	switchTest()
}

func ifTest() {
	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := computedValue(); x > 10 { //x 的作用域只在 if 块内
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}
	// println(x) x 未定义
}

func computedValue() int {
	return 3 * 5
}

func gotoTest() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签,标签名是大小写敏感的。
	println(i)
	if i < 1000 {
		i++
		goto Here //跳转到Here去
		//break Here
	}
}

func forTest() {
	sum := 0
	for index := 1; index <= 100; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)
}

func whileTest() {
	sum := 1
Loop:
	for sum < 100 {
		if sum*2 > 100 {
			break Loop
		}
		sum += sum
		println(sum)
	}
	fmt.Println("sum is equal to ", sum)
}

func switchTest() {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4") //如果没有fallthrough，默认会自动break
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8") // break
	default:
		fmt.Println("default case")
	}
}
