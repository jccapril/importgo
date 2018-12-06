package main

import (
	"math"
	"fmt"
)

// 接口类型就是一种抽象类型，是方法的几何，其他类型实现了这些方法就是实现这个接口。

// 简单示例打印不同几何图形的面积和周长

type tmp interface {
	area() float32
}

type geometry interface {
	tmp	// 内嵌接口 geometry就会拥有tmp接口所定义的方法方法
	//area() float32
	perim() float32
}

type rect struct {
	length, width float32
}

func (r rect)area() float32  {
	return r.length * r.width
}
func (r rect)perim() float32 {
	return 2 * (r.width + r.length)
}

type circle struct {
	radius float32
}

func (c circle) area() float32  {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32  {
	return 2 * math.Pi * c.radius
}

func show(name string, param interface{}) {
	switch param.(type) {

	case geometry:
		fmt.Printf("area of %v is %v \n",name, param.(geometry).area())
		fmt.Printf("perim of %v is %v \n",name, param.(geometry).perim())
	default:
		fmt.Println("wrong type!")

	}


}


func main()  {


	rec := rect{length:1,width:2}

	show("rect", rec)

	cir := circle{1}
	show("circle", cir)

	show("string", "string param")



}
