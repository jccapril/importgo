package main

import (
	"fmt"
	"errors"
)


// 程序首先运行panic，出现故障，此时跳转到包含recover()的defer函数执行，recover捕获panic，
// 此时panic就不继续传递，但是recover之后，程序并不会返回panic那个点继续执行以后的动作
// 而是在recover这个点继续执行以后的动作，即执行上面的defer函数

func recoverFunc()  {


	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	panic(errors.New("this is test for recover"))
	fmt.Println(2)

}

func main()  {

	recoverFunc()

}
