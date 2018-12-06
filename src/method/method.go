package main

import "fmt"

type T struct {
	value int
}

// 方法（method）的声明和函数很相似, 只不过它必须指定接收者
// 接受者不是任意类型，它只能为用关键字 type 定义的类型 （例如自定义类型，结构体）
// 接受者定义的方法吗不能重复，也不能和字段重复
func (t T) F()  {}

// 方法的接受者可以同时为值或者指针
func (*T)  N(){}

func (m T) StayTheSame()  {
	m.value = 3
}

func (m *T) Update()  {
	m.value = 3
}



func main()  {

	t := T{}
	t.F()
	t.N()

	// 无论值类型T还是指针类型&T都可以同时访问F和N方法
	t1 := &T{} // 指针类型
	t1.F()
	t1.N()


	// 值和指针作为接受者的区别
	// 值作为接受者（T）不会修改结构体值，而指针 *T 可以修改
	m := T{0}
	fmt.Println(m) // {0}

	m.StayTheSame()
	fmt.Println(m) // {0}

	m.Update()
	fmt.Println(m) // {3}



}

