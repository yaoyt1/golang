// 实现计算一个数字的各个位数之和如：123=6
// 采用工作池方式
package main

import (
	"fmt"
	"math/rand"
)

type job struct {
	number int
	id     int
}

type result struct {
	job *job
	sum int
}

func main() {
	j := make(chan *job, 1000)
	r := make(chan *result, 1000)

	//调用消费者
	workpool(128, j, r)
	go printResult(r)

	//调用生产者
	producer(j)
}

//calc 计算函数
func calc(job *job, res chan *result) {
	var sum int
	number := job.number
	for number != 0 {
		tmp := number % 10
		sum += tmp
		number /= 10
	}
	r := &result{
		job: job,
		sum: sum,
	}
	res <- r
}

//消费者work不断从chan管道 中获取所需参数
func work(job chan *job, res chan *result) {
	for j := range job {
		calc(j, res)
	}
}

//workpool 这里不断开启指定个数的goroutine
func workpool(workNum int, job chan *job, res chan *result) {
	for i := 0; i < workNum; i++ {
		go work(job, res)
	}
}

//printResult打印结果
func printResult(result chan *result) {
	for r := range result {
		fmt.Printf("结果：id=%d,number=%d,计算结果=%d\n", r.job.id, r.job.number, r.sum)
	}
}

//producer 生成者
func producer(jobChan chan *job) {
	var id int = 0
	for {
		id++
		number := rand.Int()
		job := &job{
			number: number,
			id:     id,
		}
		jobChan <- job
	}
}
