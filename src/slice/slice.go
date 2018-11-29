package main

import "fmt"

func main(){

	//// 数组
	//arr := [3]int{1, 2, 3}
	//fmt.Println(arr)
	//
	//// 切片
	//sliceA := []int{1, 2, 3}
	//printHelper("sliceA", sliceA)
	//
	//// use make to initialize slice
	//sliceB := make([]int, 5)
	//printHelper("sliceB", sliceB)
	//
	//sliceC := make([]int, 5, 10)  // len=5, cap=10
	//printHelper("sliceC", sliceC)
	//
	//// 使用数组初始化切片
	//sliceD := arr[0:3]
	//printHelper("sliceD", sliceD)
	//
	//// 使用切片初始化另一个切片
	//sliceE := sliceD[:]
	//printHelper("sliceE", sliceE)
	//
	//
	//sliceF := []int{1, 2, 3}
	//printHelper("sliceF", sliceF)
	//
	//sliceF = append(sliceF,4)
	//printHelper("sliceF", sliceF)

	slice := []int{1, 2, 3, 4}
	//newSlice := slice[:]
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Println("before modifying")
	printHelper("slice", slice)
	printHelper("newSlice", newSlice)

	slice[0] = 5
	fmt.Println("after modifying")
	printHelper("slice", slice)
	printHelper("newSlice", newSlice)
}


func printHelper(name string, slice []int) {
	fmt.Println(name, ":", slice)

	fmt.Printf("len of %v: %v\n", name, len(slice))

	fmt.Printf("cap of %v: %v\n", name, cap(slice))

	fmt.Println()
}
