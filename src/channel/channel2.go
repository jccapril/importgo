package main

import "fmt"

func worker(done chan bool)  {
	fmt.Println("start working...")
	done <- true
	fmt.Println("end working...")
}

func main()  {


	/*
		一个经典的例子如下，main函数中起了一个goroutine，通过非缓冲队列的使用，能够保证在goroutine执行
		结束之前 main 函数不会提前退出
	 */
	done := make(chan bool, 1)

	go worker(done)

	<-done


}
