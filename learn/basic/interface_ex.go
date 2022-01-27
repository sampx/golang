package basic

import (
	"fmt"
	"strconv"

	. "github.com/sampx/golang/utils/commonutil"
)

//Men 定义接口：人
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

//YoungChap 定义接口：年轻人
type YoungChap interface {
	Men
	BorrowMoney(amount float32)
}

//ElderlyGent 定义接口：中年人
type ElderlyGent interface {
	Men
	SpendSalary(amount float32)
}

//Human 定义类：人
type Human struct {
	name   string
	age    int
	weight int
	phone  string
}

//SayHi 打招呼,实现接口：人
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Sing 唱歌，实现接口：人， 实现接口：年轻人
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Guzzle 打瞌睡，实现接口：人
func (h *Human) Guzzle(beerStein string) {
	fmt.Printf("%s is Guzzle Guzzle Guzzle, %s ...\n", h.name, beerStein)
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h *Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

//Student 定义类：学生，实现接口：人和年轻人
type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
	phone      string
	school     string
	loan       float32
}

//Skills 技能列表
type Skills []string

//SayHi 重写Human的method,实现接口：人
func (s *Student) SayHi() {
	fmt.Printf("Hi, I am %s, I'm in school %s, you can call me on %s\n", s.name,
		s.speciality, s.phone) //Yes you can split into 2 lines here.
}

//BorrowMoney Student实现BorrowMoney方法，实现接口：年轻人
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee 定义类：员工，实现接口：人和中年人
type Employee struct {
	Human   //匿名字段
	company string
	phone   string
	money   float32
}

//SayHi 重写Human的method,实现接口：人
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

//SpendSalary Employee实现SpendSalary方法,实现接口：中年人
//1.必须传递指针引用对象，否则传递的是对象复制的值，实际对象不会改变
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
	println(e.name, "I have money:", e.money)
}

func ComplexTest() {
	Mike := Student{Human: Human{name: "Mike", age: 25, weight: 100, phone: "111"}, speciality: "PE", phone: "123", loan: 0.2}
	Sam := Employee{Human: Human{name: "Sam", age: 41, weight: 120, phone: "000"}, company: "CU", phone: "453", money: 200000}
	lucy := Employee{Human: Human{name: "Lucy", age: 68, weight: 100, phone: "000"}, company: "CU", phone: "666", money: 500000}
	Mike.Guzzle("Ha ha")
	Sam.Guzzle("Wo wo")
	//定义Men类型接口变量
	var imen ElderlyGent
	imen = &lucy //必须传递 lucy 的内存地址（见上1），否则传递的是值，方法无法改变对象值,见下2
	imen.SayHi()
	imen.Sing("Lucy rain")
	println(lucy.money) //2 更改salary之前余额
	imen.SpendSalary(425)
	println(lucy.money) //2 更改salary之后余额，余额已更改
	imen.Guzzle("lolo")
	imen = &Sam
	imen.SayHi()
	imen.Sing("Sam rain")
	imen.SpendSalary(800)
	imen.Guzzle("Ho ho")
	fmt.Println(imen) //实现了Stringer接口String()，可以直接打印
}

//type any interface{} // 空interface不包含任何的method,因此，可以存储任意类型的数值

func blankInterfaceTest(args ...Any) {
	var a Any // 空interface不包含任何的method,因此，a可以存储任意类型的数值
	var i = 5
	s := "Hello world"
	a = i
	fmt.Printf("a is %d\n", a)
	a = s
	fmt.Printf("a is %s\n", a)
	PrintAny(args)
}

type listAny []Any

func InterfaceTest() {
	//complexTest()
	l := make(listAny, 3)
	l[0] = 1       //an int
	l[1] = "Hello" //a string
	l[2] = Student{Human: Human{name: "Mike", age: 25, weight: 100, phone: "111"}, speciality: "PE", phone: "123", loan: 0.2}
	blankInterfaceTest("2324", 111, l)
}
