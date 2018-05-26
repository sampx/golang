package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(10 * time.Millisecond)
		runtime.Gosched()
	}
}

func sum(a []int, c chan int) {
	total := 0
	fmt.Println(a)
	for _, v := range a {
		total += v
	}
	time.Sleep(1000 * time.Millisecond)
	c <- total // send total to c
	println("sum done")
}

func sum1(a []int, c chan int) {
	total := 0
	fmt.Println(a)
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
	println("sum1 done")
}

func fibonacci(n int, c chan int) {
	println("in fibonacci")
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		print("B", i+1, " ", x, " ", y, " ")
		println("")
		time.Sleep(10 * time.Millisecond)
		x, y = y, x+y
	}
	println("B finish")
}

func fibonacci1(n int, c chan int) {
	println("in fibonacci 1")
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		print("A", i+1, " ", x, " ", y, " ")
		println("")
		time.Sleep(20 * time.Millisecond)
		x, y = y, x+y
	}
	println("A finish")
	close(c) //用range读取必须手动关闭
}

func buffChanTest() {
	c := make(chan int, 10) //在10个之内是非阻塞写入的
	go fibonacci1(5, c)

	go fibonacci(5, c)
	time.Sleep(2000 * time.Millisecond)
	//close(c)
	for i := range c { //range c 能够不断的读取channel里面的数据
		fmt.Print(i)
	}

}

func chanTest() {
	a := []int{7, 2, 8, -9, 4, 0, 9}

	c := make(chan int) //无buffer，会阻塞写入
	//拆成两部分，各自计算
	go sum(a[:len(a)/2], c) //sleep 1秒
	// x := <-c                //读取c会阻塞
	// println(x)
	//time.Sleep(2000 * time.Millisecond)
	go sum1(a[len(a)/2:], c) //被阻塞无法写入c
	time.Sleep(1 * time.Millisecond)
	println("waitting...")
	x, y := <-c, <-c // 读取c会阻塞,不读取的话，sum函数不执行写入chan
	fmt.Println(x, y, x+y)
}

func routineTest() {
	go say("world")
	go say("hello")
	go say("Sam")
	go say("Lucy")
	say("Ella")
}

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectTest(v bool) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func selectDefaultTest() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	//Out:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
			//break Out
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: //有default就不会阻塞
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	//routineTest()
	//chanTest()
	//buffChanTest()
	//selectTest(true)
	selectDefaultTest()
	fmt.Println(runtime.NumCPU())
}
