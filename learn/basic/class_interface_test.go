package basic

import (
	"fmt"
	"strconv"
	"testing"
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
	loan       float32
}

//Skills 技能列表 string slice
type Skills []string

//方法重载
func TestMethodOverload(t *testing.T) {
	mark := Student{Human: Human{name: "Mark", age: 25, weight: 100, phone: "111"}, speciality: "PE", phone: "123"}
	sam := Employee{Human: Human{name: "Sam", age: 41, weight: 120, phone: "000"}, company: "CU", phone: "453"}

	mark.Human.SayHi()
	sam.Human.SayHi()

	mark.SayHi()
	sam.SayHi()
}

//属性重写
func TestAttrOverride(t *testing.T) {
	// 初始化学生Jane
	jane := Student{
		Human: Human{
			name: "Jane", age: 35, weight: 100, phone: "123",
		},
		speciality: "Biology", int: 5, phone: "456",
	}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("Her phoneNumber is ", jane.phone)             //Student里面的phone会覆盖Human中的phone，重写
	fmt.Println("Her Human phoneNumber is ", jane.Human.phone) //可以明确访问里层Human中的phone
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}

func TestAdv(t *testing.T) {
	// 我们初始化一个学生
	mark := Student{
		Human: Human{
			"Mark", 25, 120, "123",
		},
		speciality: "Computer Science",
	}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)

	mark.Human = Human{name: "Marcus", age: 55, weight: 220}
	mark.age++
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
}

type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

func TestSimple(t *testing.T) {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */
	var ptr *Books  //申明一个指向books的指针

	/* Book1 属性赋值 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookID = 6495407

	/* Book 2 key 赋值 */
	Book2 = Books{title: "Python 教程", author: "www.runoob.com", subject: "Python 语言教程", bookID: 6495700}

	//Book3 顺序赋值
	Book3 := Books{"Java 教程", "www.runoob.com", "Java 语言教程", 6495788}

	//ptr指向Book1
	ptr = &Book1
	fmt.Println(ptr)

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 bookID : %d\n", Book1.bookID)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 bookID : %d\n", Book2.bookID)

	fmt.Println(Book3)
}

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

func TestComplex(t *testing.T) {
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

func TestPrintInterface(t *testing.T) {
	var sam Men
	sam = &Human{name: "Sam", age: 41, weight: 120, phone: "000"}
	l := make([]interface{}, 3)
	l[0] = 1       //an int
	l[1] = "Hello" //a string
	l[2] = sam
	fmt.Println(l)
}
