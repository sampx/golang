package basic

import (
	"fmt"
	"testing"
)

// 全局非导出变量声明,未使用，有警告但不会报错
const y1 = 70

var (
	x1       int = 123 // 包级变量
	isActive bool
)

func TestVarBasic(t *testing.T) {
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

func TestVarConst(t *testing.T) {
	//常量子表达式的顺序有可能影响到最终的估值结果
	//子表达式3/2在编译阶段被估值,3/2的估值结果为1(int)
	//表达式1.2 + 1,类型被视为float64
	//【特别注意】如果希望得到期望的浮点数运算结果，常量表达式应该写成: 3.0/2
	var x = 1.2 + 3/2
	fmt.Println(x)
}

func TestVarScope(t *testing.T) {
	// 此x变量遮挡了包级变量x1。
	var x1 = true

	// 一个内嵌代码块。
	{
		// 这里，左边的x1和y1均为新声明的变量。右边的x1为外层声明的bool变量。右边的y1为包级变量。
		// 在此内层代码块中，从此开始， 刚声明的x1和y1将遮挡外层声明x1和y1。
		x1, y1 := x1, y1-10

		x1, z := !x1, y1/10 // z是一个新声明的变量。
		// x1和y1是上一句中声明的变量。
		println(x1, y1, z) // false 60 6
	}
	println(x1) // true
	println(y1) // 70 （包级变量y1从未修改）

	// z的作用域仅限于上面的最内层代码块
	//println(z) // error: z未定义。
}
