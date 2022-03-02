package basic

import "testing"

// 下面这些新定义的类型和它们的源类型都是基本类型，但它们属于不同的类型
type (
	MyInt int
	Text  string
)

// 下面这些新定义的类型和它们的源类型都是组合类型。
type IntPtr *int
type Book struct {
	author, title string
	pages         int
}
type Convert func(in0 int, in1 bool) (out0 int, out1 string)
type StringArray [5]string
type StringSlice []string

func TestTypeFunc(t *testing.T) {
	// 这三个新定义的类型名称只能在此函数内使用。
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface{ Read([]byte) int }
	t.Logf("PersonAge type:%T", PersonAge{}) //PersonAge type:basic.PersonAge
}

// 从Go 1.9开始，可以用 = 创建类型别名，它们表示同一个类型
// 类型别名不能称为一个新类型
type (
	Name = string
	Age  = int
)

// 类型别名可以被声明在函数体内
func TestTypeFunc1(t *testing.T) {
	// 只能在此函数内使用。
	type table = map[string]int
	type Table = map[Name]Age
	t.Logf("table type:%T", Table{}) //table type:map[string]int
}

//定义类型和非定义类型
type A []string //定义类型, 也可称作有名类型
type B = A      //定义类型
// 非定义类型一定是一个组合类型. WHY?
type C = []string //非定义类型, 也可称作无名类型(这个名字有歧义,所以go1.9之后删除了这个名词)

// 这四个类型的底层类型均为内置类型int。
type (
	MyInt1 int
	Age1   MyInt1
)

// 下面这三个新声明的类型的底层类型各不相同。
type (
	IntSlice   []int   // 底层类型为[]int
	MyIntSlice []MyInt // 底层类型为[]MyInt
	AgeSlice   []Age   // 底层类型为[]Age
)

// 类型[]Age、Ages和AgeSlice的底层类型均为[]Age。
type Ages AgeSlice

// TestTypeTrans 类型推断,以及显示和隐式类型转换
//一个类型确定数字型常量所表示的值是不能溢出它的类型的表示范围的。
//一个类型不确定数字型常量所表示的值是可以溢出它的默认类型的表示范围的。
//当一个类型不确定数字常量值溢出它的默认类型的表示范围时，此数值不会被截断（亦即回绕）。
//将一个非常量数字值转换为其它数字类型时，此非常量数字值可以溢出转化结果的类型。
//在此转换中，当溢出发生时，转化结果为此非常量数字值的截断（亦即回绕）表示。
func TestTypeTrans(t *testing.T) {
	var a, b uint8 = 255, 1 // 类型确定的变量
	var c = a + b           // c==0。a+b是一个非常量表达式，结果中溢出的高位比特将被截断舍弃
	var d = a << b          // d == 254。同样，结果中溢出的高位比特将被截断舍弃
	t.Log(c, d)

	// 结果为类型不确定常量，允许溢出其默认类型。
	const X = 0x1FFFFFFFF * 0x1FFFFFFFF // 没问题，尽管X溢出
	const R = 'a' + 0x7FFFFFFF          // 没问题，尽管R溢出

	// 运算结果或者转换结果为类型确定常量
	//var e = X // error: X溢出int。
	//var h = R // error: R溢出rune。
	//const Y = 128 - int8(1)  // error: 128溢出int8。
	//const Z = uint8(255) + 1 // error: 256溢出uint8。
	//t.Log(e, h, Z)

	const a1 = -1.23 //类型不确定值，判定为字面量默认类型：float64
	var b1 = a1      //类型被推断为内置类型float64
	//var x1  = int32(a1) // error: 常量1.23不能被截断舍入到一个整数
	//var y1 int32 = b1 // error: float64类型值不能被隐式转换到int32。
	var z1 = -1        //类型推断为int32。
	var z2 = int32(b1) // float64显式转换成 int32，z的小数部分将被舍弃

	const k int16 = 255 //类型确定值: 类型为 int16，不可溢出
	var n = k           // 变量n的类型将被推断为int16。
	//var f1 = uint8(k + 1) // error: 常量256溢出了uint8。
	//var g uint8 = n + 1  // error: int16值不能隐式转换为uint8。
	// ok: h1 == 0，变量h的类型为uint8。
	// (n+1)溢出uint8，所以只有低8位
	// bits（都为0）被保留。
	var h1 = uint8(n + 1)

	_, _, _ = z1, z2, h1

}
