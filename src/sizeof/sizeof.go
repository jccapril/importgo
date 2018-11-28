package main

import (
	"fmt"
	"unsafe"
)

func  main()  {

	var (
		a int = 64
		a32bit int32 = 64
		a64bit int64 = 64
	)

	fmt.Println(a32bit == int32(a64bit))

	// 与操作系统有关，32位操作系统输出为4 64位操作系统输出为8
	fmt.Println(unsafe.Sizeof(a))

	// 32bit 输出为4
	fmt.Println(unsafe.Sizeof(a32bit))

	// 64bit 输出为8
	fmt.Println(unsafe.Sizeof(a64bit))


	var str = "Hello，世界"

    b := str[0] // b是一个uint8类型  var b uint8 = str[0]

	fmt.Println(b) // 72 ASCII码大写字母H是72，小写字母h是104 相差32

	fmt.Println(string(b)) // 输出 H

	fmt.Println(len(str)) // 输出14，输出有多少个字节

	fmt.Println(len([]rune(str))) // 输出8，输出有多少个字符




	// 这种遍历方式char是rune  int32的别名 表示一个符号
	for _, char := range str {

		//fmt.Println(string(char))
		fmt.Printf("%T", char)

	}

	fmt.Print("\n")
	// 这种遍历方式char是byte  uint8的别名 表示一个字节
	for i := 0; i <	len(str) - 1; i++ {

		char := str[i]

		fmt.Printf("%T", char)

	}


}
