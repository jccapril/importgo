package main

import (
	"fmt"
	"time"
)

// 单向通道
// 单项通道即限定了该channel只能接收或者发送数据，单向通道经常作为函数的参数
func receive(receiver chan<- string,msg string)  {

	receiver <- msg

}

func send(sender <-chan string,receiver chan<- string) {
	msg := <-sender
	receiver <- msg
}


// 需要注意的是，在变量声明中是不应该出现单向通道的，因为通道本来就是为了通信而生
// 只能接收或者发送数据是没有意义的


func main()  {

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	receive(ch1,"pass message")

	send(ch1,ch2)
	fmt.Println(<-ch2)


	// channel 遍历和关闭
	/*
	close()函数可以用于关闭 channel，关闭后的 channel 中如果有缓冲数据，依然可以读取，但是无法
	在发送数据给已经关闭的channel
	*/
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch<-i
	}
	close(ch)
	res := 0
	for v := range ch {
		res += v
	}
	fmt.Println(res)


	// select 语句
	/*
		select 专门用于通道发送和接收操作，看起来和 switch 很相似，但是进行选择和判断的方法完全不同。
		在下述例子中，通过select的使用，保证了worker中的事务可以执行完毕后才能退出main函数
	*/

	chStr := make(chan string)
	chInt := make(chan int)

	go strWorkder(chStr)
	go intWorker(chInt)

	for i := 0; i < 2; i++  {
		select {
			case <-chStr:
			fmt.Println("get value from strWorker")

			case <-chInt:
			fmt.Println("get value from intWorker")
		}
	}


}

func strWorkder(ch chan string)  {
	time.Sleep(1 * time.Second)
	fmt.Println("do something with strWorker...")
	ch <- "str"
}

func intWorker(ch chan int)  {
	time.Sleep(2 * time.Second)
	fmt.Println("do something with intWorker...")
	ch <- 1
}



