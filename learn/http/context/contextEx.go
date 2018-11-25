package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func checkCTX(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		//logg.Println(ctx)
		select {
		case s := <-ctx.Done():
			logg.Println("ctx canceled:", s)
			return
		default:
			logg.Printf("work")
		}
	}
}

func cancelContextTest() {
	ctx, cancel := context.WithCancel(context.Background())
	go checkCTX(ctx)

	//10秒钟后cancel context
	time.Sleep(3 * time.Second)
	logg.Printf("canceling...")
	cancel()
	logg.Printf("done.")
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	//cancelContextTest()
	timeoutContextTest()
}

func timeoutContextTest() {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))

	go checkCTX(ctx)

	time.Sleep(4 * time.Second)
	logg.Printf("canceling...")
	cancel()
	logg.Printf("done.")
}
