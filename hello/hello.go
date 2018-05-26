package main

import (
	f "fmt" //去GOROOT环境变量指定目录下去加载该模块,f是别名
	//绝对路径，加载gopath/src下的模块
	m "github.com/sampx/golang/utils/mathutil"
	s "github.com/sampx/golang/utils/stringutil"
)

func main() {
	f.Printf(s.Reverse("\n!oG ,olleH"))
	f.Printf("Sqrt(2) = %v\n", m.Sqrt(2))
}
