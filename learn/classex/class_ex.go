package main

import (
	"fmt"
)

//Human class
type Human struct {
	name   string
	age    int
	weight int
	phone  string
}

//Student class
type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
	phone      string
}

//Employee class
type Employee struct {
	Human   //匿名字段
	company string
	phone   string
}

//Skills list
type Skills []string

//SayHi 在human上面定义了一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//SayHi 重写Human的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

//SayHi 重写Human的method
func (s *Student) SayHi() {
	fmt.Printf("Hi, I am %s, I'm in school %s, you can call me on %s\n", s.name,
		s.speciality, s.phone) //Yes you can split into 2 lines here.
}

/*
Books Class as struct
*/
type Books struct {
	title   string
	author  string
	subject string
	bookID  int
}

func main() {
	//simpleTest()
	//advTest()
	//attrOverideTest()
	methodOverloadTest()

}

//方法重载
func methodOverloadTest() {
	mark := Student{Human: Human{name: "Mark", age: 25, weight: 100, phone: "111"}, speciality: "PE", phone: "123"}
	sam := Employee{Human: Human{name: "Sam", age: 41, weight: 120, phone: "000"}, company: "CU", phone: "453"}

	mark.Human.SayHi()
	sam.Human.SayHi()

	mark.SayHi()
	sam.SayHi()
}

//属性重写
func attrOverideTest() {
	// 初始化学生Jane
	jane := Student{Human: Human{name: "Jane", age: 35, weight: 100, phone: "123"}, speciality: "Biology", int: 5, phone: "456"}
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

func advTest() {
	// 我们初始化一个学生
	mark := Student{Human: Human{"Mark", 25, 120, "123"}, speciality: "Computer Science"}

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
	mark.Human.age++
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
}

func simpleTest() {
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
