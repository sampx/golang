package main

import (
	"errors"
	"fmt"
)

const (
	//Pi 可输出常量
	Pi = 3.1415926
	//MaxThread 可输出常量
	MaxThread = 10
)

var isActive bool // 全局变量声明

func main() {
	varTest()
	constTest()
	stringTest()
	errTest()
}

func constTest() {
	const (
		i = 10000
		s = i //const s 未使用，不会报错
	)
	prefix := "astaxie_"
	println(prefix, i)
}

func varTest() {
	var available bool // 一般声明, 默认false
	valid := false     // 简短声明
	available = true   // 赋值操作
	println(available, valid)
	//rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。
	//其中rune是int32的别称，byte是uint8的别称。
	a := 3 // int
	var b int32 = 4
	c := a + int(b) //int虽然也是32位，但int和int32是不同的数据类型,因此b要做强制类型转换
	println(c)
	v1, v2, v3 := 1, 2.0, 3 //简短声明,只能用在函数内部
	_, d := 34, 35          //将值35赋予b，丢弃34
	enabled, disabled := true, false
	println(v1, enabled, v2, disabled, v3, d) //v2 float64，go没有float类型, float32和float64两种，默认是float64
	var e complex64 = 5 + 5i
	fmt.Printf("Value is: %v\n", e) //output: (5+5i)
}

func stringTest() {
	var frenchHello string       //声明变量为字符串的一般方法
	japaneseHello := "Konichiwa" //声明 + 赋值
	var emptyString = ""         // 声明了一个字符串变量，初始化为空字符串
	println(frenchHello == emptyString)

	frenchHello = "Bonjour"                //常规赋值
	no, yes, maybe := "no", "yes", "maybe" // 简短声明，同时声明多个变量

	println(frenchHello, japaneseHello)
	println(emptyString, no, yes, maybe)
	emptyString = "str"
	// emptyString[0] = 'c' //string是不可变的
	bEmptyString := []byte(emptyString) //转变成byte数组后可以改变
	bEmptyString[0] = 'c'
	emptyString = string(bEmptyString)
	fmt.Printf("%s\n", emptyString)
	//字符串连接
	fmt.Printf("%s\n", emptyString+maybe)
	//字符串切片
	emptyString = "a" + emptyString[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", emptyString)
	//多行字符串
	mlineString := `hello
	                      world`
	fmt.Printf("%s\n", mlineString)
}

func errTest() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
