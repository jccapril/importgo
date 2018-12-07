package main

import (
	"log"
	"time"
)

// 锁
// 共享内存
// 在两个goroutine中我们都可以访问 name 这个变量，当修改它后，在不同的 goroutine 中都可以同时获取到最新的值
var name string

func main()  {


	name = "小明"

	go printName()
	go printName()

	time.Sleep(time.Second)

	name = "小红"

	go printName()
	go printName()

	time.Sleep(time.Second)

}

func printName()  {
	log.Println("name is ", name)
}