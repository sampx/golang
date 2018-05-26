package main

import (
	"fmt"
	"os"
)

var user string

func init() {
	defer errHandler() //执行完getUserEnv()后，延迟执行本函数，
	getUserEnv()
	return
}

func errHandler() { //检查是否已进入panic，并进行恢复recover()
	if x := recover(); x != nil {
		user = "default"
		fmt.Println(x)
		println("recovered")
	}
}

func getUserEnv() {
	user = os.Getenv("USER1")
	if user == "" {
		panic("no value for $USER") // 进入Panic
		//println("recover？") //无法到达这里
	}
	println("getUserEnv done")
}

func main() {
	println("user:", user)
}
