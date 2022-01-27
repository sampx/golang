package collectionex

import "fmt"

func ArrTest() {
	var arr [10]int                                 // 声明了一个int类型的数组
	arr[0] = 42                                     // 数组下标是从0开始的
	arr[1] = 13                                     // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])  //返回未赋值的最后一个元素，默认0

	//[3]int与[4]int是不同的类型
	a := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b := [4]int{1, 2, 3}   // 声明了一个长度为4的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Printf("%d,%d,%d\n", a, b, c)
}

func DArrTest() {
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("%d\n", doubleArray)
	println(doubleArray == easyArray)
}
