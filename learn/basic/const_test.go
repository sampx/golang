package basic

import (
	"fmt"
	"testing"
)

//包级常量定义
const (
	//Pi 可输出常量 未使用，有警告但不会报错
	Pi = 3.1415926
	//MaxThread 可输出常量
	MaxThread = 10 //未使用，有警告但不会报错
)

func TestConstBasic(t *testing.T) {
	//函数级常量
	const (
		i = 10000
		s = i //const s 未使用，有警告但不会报错
	)
	prefix := "astatine_"
	println(prefix, i)
	println(MaxThread, Pi, isActive)
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

func TestConstOverflow(t *testing.T) {
	const (
		// 类型不确定常量,所表示的值可以溢出其默认类型
		// 常量标识符将在编译的时候被其绑定的字面量所替代
		n = 1 << 64          // 默认类型为int
		r = 'a' + 0x7FFFFFFF // 默认类型为rune
		x = 2e+308           // 默认类型为float64
		// 类型确定常量，编译器不允许溢出
		//n1 int = 1 << 64 // error: 溢出int
		//r1 rune    = 'a' + 0x7FFFFFFF // error: 溢出rune
		//x1 float64 = 2e+308 // error: 溢出float64
	)
	//t.Log(n1, x1/2)
	t.Log(n>>2, r-0x7FFFFFFF, x/2)

}
