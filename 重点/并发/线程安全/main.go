package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int32

func main() {
	//syncAdd()
	//noSyncAdd()
	//readWriteAdd()
	yuanziAdd()
	fmt.Println(x)
}

//没有使用sync互斥锁
func noSyncAdd() {
	for i := 0; i < 100; i++ {
		go func() { x += 1 }()
	}
}

//使用互斥锁
func syncAdd() {
	var wg sync.WaitGroup
	var m sync.Mutex //互斥锁

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(&wg, &m)
	}
	wg.Wait()
}
func add(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}

//使用读写锁
func readWriteAdd() {
	var wg sync.WaitGroup
	var rw sync.RWMutex //读写锁

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go rwAdd(&wg, &rw)
	}
	wg.Wait()
}
func rwAdd(wg *sync.WaitGroup, rw *sync.RWMutex) {
	rw.Lock()
	x += 1
	rw.Unlock()
	wg.Done()
}

//原子锁
func yuanziAdd() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { atomic.AddInt32(&x, 1); wg.Done() }()
	}
	wg.Wait()
}
