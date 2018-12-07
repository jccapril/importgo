package main

import (
	"sync"
	"time"
	"fmt"
	"sync/atomic"
)

// 原子操作
/*
	atomic 提供的原子操作能够确保任一时刻只有一个goroutine对变量进行操作，
	善用 atomic 能够避免程序中出现大量的所操作
	atomic常见操作有：
	1、增减
	2、载入
	3、比较并交换
	4、交换
	5、存储
 */

/*
	增减操作
	atomic 包中提供了如下以Add为前缀的增减操作

		func AddInt32(addr *int32, delta int32)(new int32)
		func AddInt64(addr *int64, delta int64)(new int64)
		func AddUint32(addr *uint32, delta uint32)(new uint32)
		func AddUint64(addr *uint64, delta uint64)(new uint64)
		func AddUintptr(addr *uintptr, delta uintptr)(new uintptr)

	需要注意的是，第一个参数必须是指针类型的值，通过指针变量可以获取被操作数在内存中地址，
	确保同一时间只有一个 goroutine 能够进行操作

 */

// 分别用"锁"和原子操作实现多个 goroutine 对同一个变量进行累加操作

func lockAddNumber() {

	var (
		mux   sync.Mutex
		total uint64
	)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				mux.Lock()
				total += 1
				mux.Unlock()
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	// 保证 total 没有在goroutine进行读写操作
	fmt.Println("The total number is ", atomic.LoadUint64(&total))

}

func atomicAddNumber() {

	var total uint64

	for i := 0; i < 10; i++ {
		go func() {
			for {
				atomic.AddUint64(&total, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("The total number is ", atomic.LoadUint64(&total));
}

func main() {

	//lockAddNumber()
	atomicAddNumber()

}


// 载入操作
/*
	atomic 包中提供了如下以Load为前缀的载入操作

		func LoadInt32(addr *int32)(val int32)
		func LoadInt64(addr *int64)(val int64)
		func LoadPointer(addr *unsafe.pointer)(val unsafe.pointer)
		func LoadUint32(addr *uint32)(val uint32)
		func LoadUint64(addr *uint64)(val uint64)
		func LoadUintptr(addr *uintptr)(val uintptr)

	载入操作能够保证原子的读变量的值，当读取的时候，任何其他 goroutine 都无法对该变量进行读写
 */


 // 比较并交换
 /*
 	该操作简称 CAS(Compare And Swap)。这类操作的前缀为 CompareAndSwap :

 		func CompareAndSwapInt32(addr *int32, old, new int32)(swapped bool)
 		func CompareAndSwapInt64(addr *int64, old, new int64)(swapped bool)
 		func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer)(swapped bool)
 		func CompareAndSwapUint32(addr *uint32, old, new uint32)(swapped bool)
 		func CompareAndSwapUint64(addr *uint64, old, new uint64)(swapped bool)
 		func CompareAndSwapUintptr(addr *uintptr, old, new uintptr)(swapped bool)

 	该操作在进行交换前首先确保变量的值未被更改， 即仍然保持参数 old 所记录的值， 满足此前提下才进行交换操作。
 	CAS的做法类似操作数据库时常见的乐观锁机制。需要注意的是，当有大量的goroutine对变量进行读写操作时，可呢个导致CAS操作
 	无法成功，这时可以利用for循环多次尝试。

  */

// 交换
/*
	此类操作的前缀为Swap

	func SwapInt32(addr *int32, new int32) (old int32)
	func SwapInt64(addr *int64, new int64) (old int64)
	func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
	func SwapUint32(addr *uint32, new uint32) (old uint32)
	func SwapUint64(addr *uint64, new uint64) (old uint64)
	func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)

	相对于CAS，明显此类操作更为暴力直接，并不管变量的旧值是否被改变，直接赋予新值然后返回被替换的值
 */


 // 存储
 /*
 	此类操作的前缀为 Store:
 	func StoreInt32(addr *int32, val int32)
 	func StoreInt64(addr *int64, val int64)
 	func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
 	func StoreUint32(addr *uint32, val uint32)
 	func StoreUint64(addr *uint64, val uint64)
	func StoreUintptr(addr *uintptr, val uintptr)

 	在原子地存储某个值的过程中，任何 goroutine 都不会进行针对同一个值的读或写操作。
  */


  // 结论
  /*
  	atomic 和锁的方式其性能没有太大区别。
	atomic 写法相较于使用锁，更简单。
   */