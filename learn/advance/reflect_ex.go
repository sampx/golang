package advance

import (
	"fmt"
	"reflect"
)

func ReflectTest() {
	var x = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	p := reflect.ValueOf(&x)
	fmt.Println("p type:", p.Type())
	e := p.Elem()
	fmt.Println("elemt:", e)
	e.SetFloat(3.6)
	fmt.Println("elemt:", e)
}
