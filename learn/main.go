package main

import (
	"log"
	"os"
	"runtime"

	"github.com/sampx/golang/learn/advance"
	"github.com/sampx/golang/learn/basic"
	"github.com/sampx/golang/learn/collectionex"
	"github.com/sampx/golang/learn/goroutineex"
	"github.com/sampx/golang/learn/http/webserverex"
)

var logger *log.Logger

func main() {
	logger = log.New(os.Stdout, "", log.Lmicroseconds|log.Lshortfile)
	run_basic_ex()
	run_class_ex()
	run_coll_ex()
	run_routine_ex()
	run_adv_ex()
}

func run_basic_ex() {
	basic.DeferTest()
	basic.IotaTest()
	basic.VarTest()
	basic.ConstTest()
	basic.StringTest()
	basic.ErrTest()
	basic.ForTest()
	basic.IfTest()
	basic.GotoTest()
	basic.WhileTest()
	basic.SwitchTest()
	basic.SimplefuncTest()
	basic.FuncTypeTest(0)
	basic.MultiValFuncTest()
	basic.PassValPeraTest()
	basic.PassPeraPtrTest()
	basic.StructTest()
}

func run_class_ex() {
	basic.CircleAreaTest()
	basic.RectangleAreaTest()
	basic.ComplexTest()
	basic.SimpleTest()
	basic.AdvTest()
	basic.AttrOverideTest()
	basic.MethodOverloadTest()
	basic.InterfaceTest()
}

func run_coll_ex() {
	collectionex.ArrTest()
	collectionex.DArrTest()
	collectionex.SliceTest()
	collectionex.AdvSliceTest()
	collectionex.SliceFilterTest()
	collectionex.SimpleMapTest()
	collectionex.AdvMapTest()
	collectionex.RangeTest()
}
func run_routine_ex() {
	logger.Println("cpus: ", runtime.NumCPU())
	goroutineex.RoutineTest()
	goroutineex.ChanTest()
	goroutineex.BuffChanTest()
	goroutineex.SelectTest(true)
	goroutineex.SelectDefaultTest()
	goroutineex.ReadWriteTest()
	goroutineex.ChanClosetest()
	goroutineex.RangeReadTest()
}

func run_adv_ex() {
	advance.ClosureTest()
	advance.PtrArrayTest()
	advance.JsonTest()
	advance.PanicTest()
	advance.ReflectTest()
	webserverex.StartHttpServer()
	webserverex.HttpContexTest()
}
