package main

import (
	"fmt"
	"time"
)

func main() {
	//basicChanel()
	//chanelInOut()
	//goroutineSync()
	//singleChanel()
	//forRangeChan()
	selectChan()
}

//basicChanel 基本使用
func basicChanel() {
	var ch chan int
	if ch == nil {
		fmt.Println("chan管道为空")
	}
	fmt.Printf("ch类型为：%T\n", ch)
}

// chan 传入传出
func chanelInOut() {
	var ch chan int
	if ch == nil {
		fmt.Println("chan管道为空")
	}
	ch = make(chan int, 1)
	ch <- 100
	fmt.Printf("ch类型为：%T\n", ch)
	data := <-ch
	fmt.Printf("ch=%v,data=%v", ch, data)
}

// chan 传入传出
func chanelInOut1() {
	var ch chan int
	if ch == nil {
		fmt.Println("chan管道为空")
	}
	ch = make(chan int)
	go func() { ch <- 100 }()
	fmt.Printf("ch类型为：%T\n", ch)
	data := <-ch
	fmt.Printf("ch=%v,data=%v", ch, data)
}

//chan实现goroutine同步
func goroutineSync() {
	var c chan bool = make(chan bool)
	go helloWord(c)
	<-c
	fmt.Println("goroutine 同步完成")
}
func helloWord(c chan bool) {
	fmt.Println("goroutine hello word ")
	c <- true
}

//chanel 的单向操作
func singleChanel() {
	c := make(chan int)
	go writeChan(c)
	readChan(c)
}

func readChan(c <-chan int) {
	<-c
	fmt.Println("读取")
}

func writeChan(c chan<- int) {
	c <- 10
	fmt.Println("写入")
}

//chanel 循环
func forRangeChan() {
	var c chan int = make(chan int)
	go func(a chan int) {
		for i := 0; i < 5; i++ {
			a <- i
		}
	}(c)

	for v := range c {
		fmt.Println(v)
	}
}

//select 操作
/*
判断chan是否是空的
判断chan是否是满的
case 随机策略
*/
func selectChan() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go server1(ch1)
	go server2(ch2)

	time.Sleep(time.Second)

	select {
	case s1 := <-ch1:
		fmt.Println("ch1", s1)
	case s1 := <-ch2:
		fmt.Println("ch2", s1)
	default:
		fmt.Println("暂时没有管道插入")
	}
}
func server1(c chan<- string) {
	//time.Sleep(time.Second * 6)
	c <- "from server1"
}
func server2(c chan<- string) {
	//time.Sleep(time.Second * 3)
	c <- "from server2"
}
