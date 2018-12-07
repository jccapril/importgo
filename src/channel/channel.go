package main

import "fmt"

// channel
// goroutine是Go中实现并发的重要机制，channel是goroutine之间进行通信的重要桥梁


func main()  {

	// 使用make函数创建channel
	//ch := make(chan int)

	// 也可以使用var声明channel
	//var ch chan int

	// channel是双向的，意味可以该channel可以发送数据，也可以接受数据
	//ch<-x // channel 接受数据 x
	//x<-ch // channel 发送数据并赋值给x
	//<-ch	// channel 发送数据，忽略接受者


	// channel buffer，channel可以定义缓冲大小
	//chInt := make(chan int) // unbuffered channel 非缓冲通道
	//chBool := make(chan bool, 0) // unbuffered channel 非缓冲通道
	//chStr := make(chan string, 2) // buffered channel 缓冲通道



	ch := make(chan string)

	// main函数是一个goroutine ，在这一个goroutine中发送了数据给非缓冲通道
	// 但是却没有另外一个goroutine从非缓冲通道里去读取数据，所以造成了阻塞或成为死锁
	//ch <-"ping"

	go func() {
		ch<-"ping"
	}()



	fmt.Println(<-ch)


	// 与非缓冲通道不同， 缓冲通道可以在同一个goroutine内接受容量范围内的数据，即便没有
	// 另外的goroutine进行读取操作
	ch1 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2

	fmt.Println(<-ch1)
	ch1 <- 3
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}
