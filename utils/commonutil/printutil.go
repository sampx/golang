package commonutil

import (
	"fmt"
	"reflect"
)

//PrintAny 打印任何类型的slice
func PrintAny(ls ...interface{}) {
	for index, element := range ls {
		//注意：element.(type),此语法不能在switch以外使用，
		//在switch之外可以使用: if value, ok := element.(int); ok { //comma ok 断言 类型
		switch value := element.(type) {
		case int, rune, int64, float32, float64, byte:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		default:
			v := reflect.ValueOf(element)
			fmt.Printf("list[%d] is of a different type, its type is %s\n", index, v.Type())
		}
	}
}
