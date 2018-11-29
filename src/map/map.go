package main

import (
	"fmt"

	//"sync"
)

func main()  {

	/*
	var aMap map[string]int  // 只定义， 此时 cMap 为nil
	fmt.Println(aMap == nil)
	//cMap["北京"] = 1 	// 报错，因为 cMap 为nil

	bMap := make(map[string]int)
	bMap["北京"] = 1

	// 制定初始容量
	cMap := make(map[string]int, 100)
	cMap["北京"] = 1

	//dMap := map[string]int{"北京", 1}

	code := cMap["北京"] // 读
	fmt.Println(code)

	code = cMap["广州"] // 读不存在的key
	fmt.Println(code)

	code, ok := cMap["广州"]
	if ok {
		fmt.Println(code)
	}else {
		fmt.Println("key not exist")
		fmt.Println(code)
	}

	delete(cMap, "北京") // 删除 key
	fmt.Println(cMap)

	eMap := map[string]int{"北京":1, "上海":2, "广州":3, "深圳":4}
	for city, code := range eMap {
		fmt.Printf("%s:%d", city, code)
		fmt.Println()
	}
	*/
	/*
	fMap := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fMap["北京"] = 1
		wg.Done()
	}()

	go func() {
		fMap["上海"] = 2
		wg.Done()
	}()

	wg.Wait()
	*/
	provinces := make(map[string]map[string]int)

	provinces["北京"] = map[string]int{
		"东城区": 1,
		"西城区": 2,
		"朝阳区": 3,
		"海淀区": 4,
	}

	fmt.Println(provinces)

}
