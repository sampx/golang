package basic

import (
	"fmt"
)

//Person struct
type Person struct {
	Name string
	Age  int
}

//实现stringer接口
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func StructTest() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, "\n", z)
}
