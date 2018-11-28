package main

import "fmt"

func main()  {



}

func printHelper(name string, arr [5]int)  {

	for i := 0; i < 5; i++  {

		fmt.Printf("%v[%v] : %v\n",name, i, arr[i])

	}

	// len获取长度
	fmt.Printf("len of %v: %v\n")

}
