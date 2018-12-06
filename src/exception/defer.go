package main

import (
	"fmt"
	"net/http"
	"os"
)

func main()  {

	// defer会在outerFunc退出之前执行打印操作，因此被defer调用的函数也称为"延迟函数"
	outerFunc()

}



func outerFunc(){
	defer fmt.Printf(" World!\n")

	fmt.Print("Hello")
}

// defer常用场景
// defer语句经常被用于处理成对的操作，如打开和关闭，连接和断开连接，加锁和释放锁。
// 恰当使用defer能够保证资源正确释放。

// 使用 defer 关闭 http 请求响应体的 Body
func closeBody(url string) error {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	// ... do more stuff ...

	return err
}


// 使用 defer 关闭文件句柄
func closeFile(filename string) error {
	f, err := os.Open(filename)
	defer  f.Close()
	// ... do more stuff ...

	return err
}



// 使用 defer 解锁
/*
func BillCustomer(c *Customer)  {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// ... do more stuff ...

	return
}
*/

// defer 使用中的注意点和例子

// 例子1
// 最终打印结果是 5 5 5 5 5
// 这是因为 defer 所调用的函数是延迟执行的。等到执行 defer 所调用的函数时，i已经是5了。
func printNumer(){
	for i := 0; i < 5; i++{
		defer func() {
			fmt.Println(i)
		}()
	}
}


// 例子2
// 最终打印结果是 4 3 2 1 0
func printNum(){
	for i := 0; i < 5; i++ {
		defer func(v int) {
			fmt.Println(v)
		}(i)
	}
}

// 例子3
// 最终返回的是4
// 因为 return 2 执行后，变量 i 赋值为2，但是随后执行了 defer 函数，i 被赋值为4，所以最终返回的是4
func testDefer()(i int){
	defer func() {
		fmt.Println(i)
		i = 4
	}()
	return 2
}