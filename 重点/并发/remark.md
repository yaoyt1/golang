# 并发编程

## 多线程
* 线程是属于操作系统进行管理， 也就是属于内核态
* 线程之间切换， 需要发生用户态和内核态的切换   
* 当系统中有大量线程的时候，系统会变卡
* 用户态的线程，支持大量的创建线程，也叫协程或gouroutine

## gouroutine
> 一个操作系统对应用户态多个gouroutine
> 同时使用多个操作系统线程
> 操作系统线程对gouroutine时多对多的关系，即m:n


## chan 
> 本质上是个容器
> var chanName chan type 

### chan 基本使用

```go
var ch chan int
if ch == nil {
	fmt.Println("chan管道为空")
}
fmt.Printf("ch类型为：%T", ch)
```

### chan 传入传出

```go
var ch chan int
if ch == nil {
	fmt.Println("chan管道为空")
}
//这里我们初始化chan 并给出长度
//如果长度没有给直接传入会报错
ch = make(chan int, 1)
ch <- 100
fmt.Printf("ch类型为：%T\n", ch)
data := <-ch
fmt.Printf("ch=%v,data=%v", ch, data)
```

```go
var ch chan int
if ch == nil {
fmt.Println("chan管道为空")
}

// 这里我们初始化并没有给出长度
// 但是我们新开了一个gouroutine 来传入，并外面有传出操作 
ch = make(chan int)
go func() { ch <- 100 }()
fmt.Printf("ch类型为：%T\n", ch)
data := <-ch
fmt.Printf("ch=%v,data=%v", ch, data)
```

### chan 阻塞
> 只传入不传出就会造成阻塞

### chan 实现 gouroutine同步
```go
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
```
### chan 单向操作
```go
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
```
### chan 关闭
```go
close(c)
```
### chan for range 循环
```go
//chanel 循环
//这里程序会死锁，因为生产者没有传入，程序一直等待，故报错了
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
```
### chan 带缓冲区
```go
var c chan int =make(chan int ,1)
```

### selcet 操作
```go
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

```


## 线程安全
### 互斥锁

### 读写锁

### 原子锁

