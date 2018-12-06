package main

import (
	"log"
	"time"
	"sync"
	"fmt"
)

/*
	并发指的是多个任务被（一个）cpu轮流切换执行，在Go语言里面主要用goroutine(协程)来实现并发
 	类似于其他语言中的线程。
 */

//func doSomething(id int)  {
//	log.Printf("before do job:(%d) \n", id)
//	time.Sleep(3 * time.Second)
//	log.Printf("after do job:(%d) \n", id)
//}

func doSomething(id int, wg *sync.WaitGroup)  {
	defer wg.Done()
	log.Printf("before do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d) \n", id)
}


func main()  {

	//go doSomething(1)
	//go doSomething(2)
	//go doSomething(3)
	//
	//time.Sleep(4 * time.Second)


	//var wg sync.WaitGroup
	//wg.Add(3)
	//go doSomething(1, &wg)
	//go doSomething(2, &wg)
	//go doSomething(3, &wg)
	//
	//
	//
	//wg.Wait()
	//
	//log.Printf("finish all jobs \n")


	/*
		运行结果是 3  3 3
		原因：
		1、所有goroutine代码片段中的i 是同一个变量，待循环结束的时候，它的值是3
		2、main()循环结束后才开始并发执行新生成的goroutine
	  */
	  /*
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	*/

	/*
		解决方案
	 */
	for i := 0; i < 3; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
	}


	time.Sleep(1 * time.Second)

}
