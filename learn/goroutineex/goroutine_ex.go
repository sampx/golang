package goroutineex

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.Lmicroseconds|log.Lshortfile)
}

func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Second)
		runtime.Gosched()
	}
}

func sum(a []int, c chan int) {
	logger.Println("sum start")
	total := 0
	for _, v := range a {
		total += v
	}
	logger.Println("sum计算比较耗时,请稍候1秒....")
	time.Sleep(1000 * time.Millisecond)
	logger.Println("sum done")
	c <- total // send total to c
}

func sum1(a []int, c chan int) {
	logger.Println("sum1 start")
	total := 0
	logger.Println("sum1计算比较耗时,请稍候500毫秒....")
	time.Sleep(500 * time.Millisecond)
	for _, v := range a {
		total += v
	}
	logger.Println("sum1 done")
	c <- total // send total to c
}

func fib(n int, c chan int) {
	logger.Println("fib start...")
	start := time.Now()
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		println("N", i+1, " ", x, " ", y, " ")
		time.Sleep(10 * time.Millisecond)
		x, y = y, x+y
	}
	end := time.Now()
	logger.Println("fib done. spend:", end.Sub(start))
	close(c) //如果用range读取,此处必须手动关闭chan,否则range会持续读取导致主线程deadlock
}

func BuffChanTest() {
	c := make(chan int, 5) //在5个之内是非阻塞,也就是说即使不读取chan c,发送方可以持续写入5个元素到缓存队列,然后就会阻塞
	logger.Println("开始fib....")
	go fib(10, c) //长度大于5,如果没有接收方读取,fib函数将执行到一半后会阻塞
	logger.Println("主线程休眠1秒...")
	time.Sleep(1000 * time.Millisecond)
	//close(c) 此处不能close,否则后续无法读取chan c
	logger.Println("主线程读取结果...")
	var rst = make([]int, 0, 10)
	for i := range c { //不断的读取channel里面的数据,此时才开始进行fib函数计算
		rst = append(rst, i)
	}
	logger.Println(rst)
	logger.Println("all done.")
}

func ChanTest() {
	a := []int{7, 2, 8, -9, 4, 0, 9}
	//拆成两部分，各自计算和,然后再汇总起来
	logger.Println("sum开始计算第一部分:", a[:len(a)/2])
	c := make(chan int)     //无buffer，会阻塞写入
	go sum(a[:len(a)/2], c) //耗时1秒
	// x := <-c          //此处读取c会阻塞,直到sum结果返回chan c,才执行下一步,因此如果想并行,此处先不要读取chan
	// logger.Println("sum 结果:", x)
	//time.Sleep(2000 * time.Millisecond) //如果此处sleep实际上也没有并行起来,因为sum的耗时小于两秒
	logger.Println("sum1开始计算第二部分:", a[len(a)/2:])
	go sum1(a[len(a)/2:], c) //耗时500毫秒
	y := <-c                 //如果不读取chan, 将不会执行sum里的运算,很智能的延迟加载,这里是先完成向chan中输出的那个值,即sum1的结果
	logger.Println("先完成运算的结果:", y)
	x := <-c //这里是先后成向chan中输出的那个值,即sum的结果
	logger.Println("后完成运算的结果:", x)
	//x = <-c //此处如果再次读取也会报错,因为chan中不可能再产生数据,goroutines会deadlock
	logger.Println(x, y, x+y)
}

func RoutineTest() {
	for index := 0; index < 10; index++ { //启动10个gorutine
		go say(strconv.Itoa(index))
	}
	say("Ella")
}

func fib1(c, quit chan int) {
	logger.Println("fib1 start...")
	x, y := 1, 1
	for { //无限循环
		select {
		case c <- x: //将x发送到chan c,如果无接收方队列,此处会阻塞.有接收方的话,此时接收方退出阻塞,可以接收到x的值
			logger.Println("set:", x)
			//time.Sleep(10 * time.Millisecond)
			x, y = y, x+y
		case <-quit: //quit chan无发送方队列,此处阻塞.当quit chan中有值被发送进来,退出循环
			logger.Println("quit")
			return
		}
	}
}

func SelectTest(v bool) {
	c := make(chan int) //创建一个无缓存的chan
	quit := make(chan int, 1)
	rst := make([]int, 0, 10)
	logger.Println("启动计算协程fib1...")
	go fib1(c, quit)

	for i := 0; i < 10; i++ { //连续读取10次,当c中没有值被发送进来,此处会阻塞
		time.Sleep(2000 * time.Millisecond)
		logger.Println("get:", i)
		rst = append(rst, <-c)
	}
	//发送退出信号
	quit <- 0

	logger.Println("result: ", rst)

	logger.Println("all done.")
}

func SelectDefaultTest() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)
	//Out:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
			//break Out
		case b := <-boom:
			fmt.Println("BOOM!", b)
			return
		default: //有default就不会阻塞,每次循环如无上述匹配就执行default
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func ReadWriteTest() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	logger.Println("start writer...")
	go write(ch1)
	logger.Println("start reader...")
	go read(ch1, ch2)
	logger.Println("wait for guit signal...")
	for {
		time.Sleep(400 * time.Millisecond)
		select { //阻塞等待quit信号
		case <-ch2:
			logger.Println("all done.")
			return
		default:
			fmt.Print(".")
		}
	}
}

func read(ch1 chan int, ch2 chan bool) {
	for {
		if v, ok := <-ch1; ok { //用ok进行测试,只要能够从发送方读取,就一直接收
			fmt.Printf("read a int is %d\n", v)
		} else {
			ch2 <- true //如果发送队列空了,发送退出信号
		}
	}
}

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}
	close(ch)
}

//1. 如果channel已经关闭,继续往它发送数据会导致panic: send on closed channel
//2. 关闭一个已经关闭的channel也会导致panic: close of closed channel
//3. channel关闭后，仍然可以从中读取以发送的数据，读取完数据后，将读取到零值，可以多次读取。
//4. 用range读取channel,如果不手动关闭channel,进程将一直阻塞在当前协程,并导致所有协程阻塞死锁
func ChanClosetest() {
	ch := make(chan int, 3)
	ch <- 3
	ch <- 2
	ch <- 1
	close(ch)
	fmt.Print(<-ch) //3
	fmt.Print(<-ch) //2
	fmt.Print(<-ch) //1
	fmt.Print(<-ch) //0
	fmt.Print(<-ch) //0
}

func RangeReadTest() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for v := range data { //data只要不close会一直阻塞在此
			println(v)
		}
		exit <- true
		println("receive over")
	}()

	data <- 1
	data <- 2
	// data <- 3
	// data <- 4
	close(data)

	println("send over")
	<-exit
}
