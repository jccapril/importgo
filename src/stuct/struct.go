package main

import "fmt"

type Student struct {
	Age int
	name string
}

// 通常结构体中的一个字段占一行，但是类型相同的字段，也可以放在同一行
//type Student struct {
//	Age int
//	Name,Address string
//}

type Tree struct {
	value int
	left, right *Tree
}

type TopStudent struct {
	 Student
}



func main()  {

	stu := Student{
		Age: 18,
		name: "name",
	}
	fmt.Println(stu)

	// 在赋值时字段名可以忽略
	fmt.Println(Student{20,"new name"})


	tree := Tree{
		value:1,
		left:&Tree{
			value:1,
			left:nil,
			right:nil,
		},
		right:&Tree{
			value:2,
			left:nil,
			right:nil,
		},
	}

	fmt.Printf(">>> %#v\n", tree)

	topStu := TopStudent{stu}
	fmt.Println(topStu)


}
