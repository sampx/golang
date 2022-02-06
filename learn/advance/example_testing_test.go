package advance

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"text/template"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello + %d", i)
		//fmt.Println(s)
	}
}

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}

func ExampleHello() {
	//运行测试时函数输出与下面 Output 注释进行比较
	fmt.Println("hello") // Output: hello

}

//测试换行输出
func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}

// 和 "Output:" 类似，但是能够以任意顺序匹配行
func ExamplePerm() {
	for _, value := range [...]int{0, 1, 2, 3, 4} {
		fmt.Println(value)
	}
	// Unordered output:
	// 4
	// 2
	// 1
	// 3
	// 0
}

type Student struct {
	Name  string
	Age   int
	Class string
	Score int
}

func BenchmarkReflect_New(b *testing.B) {
	var s *Student
	sv := reflect.TypeOf(Student{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv)
		s, _ = sn.Interface().(*Student)
	}
	_ = s
}
func BenchmarkDirect_New(b *testing.B) {
	var s *Student
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = new(Student)
	}
	_ = s
}
func BenchmarkReflect_Set(b *testing.B) {
	var s *Student
	sv := reflect.TypeOf(Student{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv)
		s = sn.Interface().(*Student)
		s.Name = "Jerry"
		s.Age = 18
		s.Class = "20005"
		s.Score = 100
	}
	_ = s
}
func BenchmarkReflect_SetFieldByName(b *testing.B) {
	sv := reflect.TypeOf(Student{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv).Elem()
		sn.FieldByName("Name").SetString("Jerry")
		sn.FieldByName("Age").SetInt(18)
		sn.FieldByName("Class").SetString("20005")
		sn.FieldByName("Score").SetInt(100)
	}
}
func BenchmarkReflect_SetFieldByIndex(b *testing.B) {
	sv := reflect.TypeOf(Student{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sn := reflect.New(sv).Elem()
		sn.Field(0).SetString("Jerry")
		sn.Field(1).SetInt(18)
		sn.Field(2).SetString("20005")
		sn.Field(3).SetInt(100)
	}
}
func BenchmarkDirect_Set(b *testing.B) {
	var s *Student
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = new(Student)
		s.Name = "Jerry"
		s.Age = 18
		s.Class = "20005"
		s.Score = 100
	}
	_ = s
}
