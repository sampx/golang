package main

import "fmt"

//slice并不是真正意义上的动态数组，而是一个引用类型。
//slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
func main() {
	sliceTest()
	advSliceTest()
}

func sliceTest() {
	slice := []byte{'a', 'b', 'c', 'd'} // 和声明array一样，只是少了长度
	fmt.Printf("%c\n", slice)
	// 声明一个含有10个元素元素类型为byte的数组
	var arr = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Printf("%c\n", arr[:3]) //arr[0:3] : a,b,c
	fmt.Printf("%c\n", arr[5:]) //arr[5:len(arr)] : 'f', 'g', 'h', 'i', 'j'
	fmt.Printf("%c\n", arr[:])  //arr[0:len(arr)] : arr的全部元素
	// 声明两个含有byte的slice
	var a, b []byte
	// a指向arr数组的第3个元素开始，并到第五个元素结束，
	a = arr[2:5] //现在a含有的元素: arr[2]、arr[3]和arr[4]
	// b是数组arr的另一个slice
	b = arr[3:] //现在b含有的元素: arr[3]到最后
	//此处添加的x把arr中该位置的f覆盖掉了，要非常小心,可以通过限制其容量来防止：a = arr[2:5:5]
	a = append(a, 'x') //append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
	//此处添加的y不会影响arr数组，因为arr数组已达到最大容量
	b = append(b, 'y')
	fmt.Printf("%c,%c,%c\n", a, b, arr)
}

func advSliceTest() {
	var numbers []int   //创建slice
	printSlice(numbers) //长度默认为0

	/* 追加空切片 */
	//内置函数append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
	//注：append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	//copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
