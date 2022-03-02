package basic

import "testing"

func TestAnonymousFunc(t *testing.T) {
	// 这个匿名函数没有输入参数，但有两个返回结果。
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // 一对小括号表示立即调用此函数。不需传递实参。

	// 下面这些匿名函数没有返回结果。
	func(a, b int) {
		println("a*a + b*b =", a*a+b*b) // a*a + b*b = 25
	}(x, y) // 立即调用并传递两个实参。

	func(x int) {
		// 形参x遮挡了外层声明的变量x。
		println("x*x + y*y =", x*x+y*y) // x*x + y*y = 32
	}(y) // 将实参y传递给形参x。

	func() {
		println("x*x + y*y =", x*x+y*y) // x*x + y*y = 25
	}() // 不需传递实参。
}

func TestConstFunction(t *testing.T) {
	// c是一个类型不确定复数常量
	// 此调用将在编译时刻被估值，其返回结果也是一个常量
	const c = complex(1.6, 3.3) // 此函数调用是一个常量表达式

	// 函数调用real(c)和imag(c)的结果都是类型
	// 不确定浮点数值。在下面这句赋值中，它们都
	// 被推断为float32类型的值。
	var a, b float32 = real(c), imag(c)

	// 变量d的类型被推断为内置类型complex64。
	// 函数调用real(d)和imag(d)的结果都是
	// 类型为float32的类型确定值。
	var d = complex(a, b)

	// 变量e的类型被推断为内置类型complex128。
	// 函数调用real(e)和imag(e)的结果都是
	// 类型为float64的类型确定值。
	var e = c

	_, _ = d, e
}
