package basic

import (
	"fmt"
)

func StringTest() {
	var frenchHello string              //声明变量为字符串的一般方法,string初始值是""
	japaneseHello := "Konichiwa"        //声明 + 赋值
	var emptyString = ""                // 声明了一个字符串变量，初始化为空字符串
	println(frenchHello == emptyString) // true

	frenchHello = "Bonjour"                //常规赋值
	no, yes, maybe := "no", "yes", "maybe" // 简短声明，同时声明多个变量

	println(frenchHello, japaneseHello)
	println(emptyString, no, yes, maybe)
	emptyString = "str"
	// emptyString[0] = 'c' //string是不可变的
	bEmptyString := []byte(emptyString) //转变成byte数组后可以改变
	bEmptyString[0] = 'c'
	emptyString = string(bEmptyString)
	fmt.Printf("%s\n", emptyString) //ctr
	//字符串连接
	fmt.Printf("%s\n", emptyString+maybe)
	//字符串切片
	emptyString = "a" + emptyString[1:] // 字符串虽不能更改，但可进行切片操作,atr
	fmt.Printf("%s\n", emptyString)
	//多行字符串
	mlineString := `hello
	                      world`
	fmt.Printf("%s\n", mlineString)
}
