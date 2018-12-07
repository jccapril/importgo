package main

import (
	"sync"
	"sync/atomic"
	"time"
	"fmt"
)

// 互斥锁和读写锁的并发性能
// sync.Mutex 是互斥锁，只有一个信号标量
// 在 Go 中还有一种读写锁 sync.RWMutex
// 如果可以分理出读和写两个互斥信号的情况，可以考虑使用它来提高读的并发性能

func main(){


	var(
		mux sync.Mutex
		state1 = map[string]int{
			"a":65,
		}
		muxTotal uint64

		rw sync.RWMutex
		state2 = map[string]int {
			"a":65,
		}
		rwTotal uint64

	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mux.Lock()
				_ = state1["a"]
				mux.Unlock()
				atomic.AddUint64(&muxTotal, 1)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				rw.RLock() // RLock对读操作加锁， Lock对写操作加锁
				_ = state2["a"]
				rw.RUnlock()
				atomic.AddUint64(&rwTotal,1)
			}

		}()
	}

	time.Sleep(time.Second)

	fmt.Println("sync.Mutex readOps is ", muxTotal)
	fmt.Println("sync.RWMutex readOps is ", rwTotal)
}
