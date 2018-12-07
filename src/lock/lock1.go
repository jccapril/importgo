package main

import (
	"sync"
	"fmt"
)

// 当我们并发对同一个切片进行写操作的时候，会出现数据不一致的问题，这就是一个典型的共享变量的问题
// 针对这个问题我们可以使用lock（锁）来修复，从而保证数据的一致性

func main()  {

	var(
		wg		sync.WaitGroup
		numbers	[]int

		mux sync.Mutex
	)

	for i :=0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mux.Lock()
			numbers = append(numbers, i)
			mux.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("The numbers is ", numbers)

}
