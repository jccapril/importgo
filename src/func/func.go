package main

import (
	"fmt"

)

func main()  {

	fmt.Println(plus(1,2))

	fmt.Println(sum(1))
	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3))


	// 匿名函数
	func(name string){
		fmt.Println(name)
	}("Hello Wolrd!")

	// 闭包
	addOne := addInt(1)
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())

	addTwo := addInt(2)
	fmt.Println(addTwo())
	fmt.Println(addTwo())
	fmt.Println(addTwo())


	// 函数作为参数
	logger(sayHello, "World")

	// 传值和传引用

	str := "蒋晨成"
	fmt.Println("before calling sendValue, str : ", str)
	sendValue(str)
	fmt.Println("after calling sendValue, str : ", str)

	fmt.Println("before calling sendAddress, str : ", str)
	sendAddress(&str)
	fmt.Println("after calling sendAddress, str: ", str)

}


// 单返回值函数

func plus(a, b int)(res int) {
	return a + b
}

// 多返回值函数
func multi()(string, int)  {
	return "name",18
}
// 命名返回值
// 被命名的返回参数的值为该类型的默认零值
// 该例子中 name 默认初始化为空字符串，height 默认初始化为0
func namedReturnValue()(name string, height int) {
	name = "xiaoming"
	height = 180
	return
}

// 参数可变函数
func sum(nums ...int)int {
	fmt.Println("len of nums is : ", len(nums))
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}




// 闭包
func addInt(n int) func() int  {
	i := 0
	return func() int {
		i += n
		return i
	}
}


// 函数作为参数
func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func logger(f func(string), name string)  {
	fmt.Println("start calling method sayHello")
	f(name)
	fmt.Println("end calling method sayHello")
}

// 传旨和传引用
func sendValue(name string)  {
	name = "jcc"
}

func sendAddress(name *string)  {
	*name = "jcc"
}
