package main

import "fmt"

type City string

type Age int

type Height int

type (
	B0 = int8
	B1 = int16
	B2 = int32
	B3 = int64
)

func main()  {


	var c City
	c = "ShangHai"

	fmt.Println(c)

	city := City("北京")
	fmt.Println("I live in",city + "上海")
	fmt.Println(len(city))

	middle := Age(12)
	if middle >= 12 {
		fmt.Println("Middle is bigger than 12")
	}

	//printAge(middle) // 会报错 必须显式转换相同类型
	printAge(int(middle))


	age := Age(12)
	height := Height(175)
	//fmt.Println(height/age) // 会报错 必须显式转换相同类型
	fmt.Println(int(height)/int(age))

}

func printAge(age int)  {
	fmt.Println("Age is ",age)
}