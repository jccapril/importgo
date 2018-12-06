package main

import (
	"fmt"
	"errors"
)

// panic
// 当程序遇到致命错误导致无法继续运行时就会触发 panic ，例如：数组越界、空指针等

// 以下代码将会发出数组越界异常
func testPanic(){
	s := []int{1,2,3}
	for i := 0; i <= 4; i++ {
		fmt.Println(s[i])
	}
}

// 上述例子中因为数组越界，出发了runtime异常，导致程序退出。在实际开发中，也可以主动调用 panic函数达到同样的效果
func panicFunc(){
	panic(errors.New("this is test for panic"))
}


func main()  {

	//panicFunc()
	//testPanic()
}
