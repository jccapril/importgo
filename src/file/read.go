package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"time"
	"sync"
	"io"
	"bufio"
)

// 读文件
// 全量读  带缓冲区读   任意位置读

// 全量读
func readFile1() {
	dat, err := ioutil.ReadFile("./read.go")
	fmt.Println(err)
	fmt.Println(string(dat))
}

// 带缓冲区读
func readFile2() {
	f, _ := os.Open("./read.go")
	defer f.Close()

	buf := make([]byte, 16)
	f.Read(buf)

	fmt.Println(string(buf))
}

// 任意位置读
// 1、Seek()和Read()

func readFile3() {

	f, _ := os.Open("./read.go")
	defer f.Close()

	buf := make([]byte, 2)

	// 第二个参数即whence参数：0表示从头开始，1表示当前位置，2表示从尾开始
	f.Seek(5, 0)
	f.Read(buf)
	fmt.Println(string(buf))

}

// 任意位置读
// 2、ReadAt()
func readFile4() {
	f, _ := os.Open("./read.go")
	defer f.Close()

	buf := make([]byte, 2)

	f.ReadAt(buf, 5)

	fmt.Println(string(buf))

}



// ReadAt 并发安全
func rrr() {
	f, _ := os.Open("./read.go")
	defer f.Close()

	for i := 0; i < 5; i++ {
		go func() {
			b1 := make([]byte, 2)
			f.ReadAt(b1, 5)
			fmt.Println(string(b1))

			//f.ReadAt(b1, 2)
			//fmt.Println(string(b1))
		}()
	}

	time.Sleep(time.Second)
}

// Seek()、Read()并发不安全
func unsafeRRR() {
	var mutex sync.Mutex
	f, _ := os.Open("./read.go")
	defer f.Close()

	for i := 0; i < 5; i++ {
		go func() {
			b1 := make([]byte, 2)
			mutex.Lock()
			f.Seek(5, 0)
			f.Read(b1)
			mutex.Unlock()
			fmt.Println(string(b1))

			//f.Seek(2, 0)
			//f.Read(b1)
			//fmt.Println(string(b1))
		}()
	}

	time.Sleep(time.Second)
}

// 使用 buf 实现 ioutil.ReadFile 类似效果
func readCustomAllFileContent() {

	f, _ := os.Open("./read.go")
	defer f.Close()

	content := make([]byte, 0)
	buf := make([]byte, 16)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		content = append(content, buf[0:n]...)
	}
	fmt.Println(string(content))
}

//使用 bufio 实现行统计：

func readLines()  {

	f, _ := os.Open("./read.go")
	defer f.Close()

	br := bufio.NewReader(f)
	totalLine := 0

	for {
		_, isPrefix, err := br.ReadLine()
		if !isPrefix {
			totalLine += 1
		}

		if err == io.EOF {
			break
		}
	}
	fmt.Println("total lines is: ", totalLine)
}

func main() {

	// 全量读
	//readFile1()

	// 带缓冲区读
	//readFile2()

	// 任意位置读
	//readFile3()

	// 任意位置读
	//readFile4()

	// 并发不安全
	//rrr()

	// 并发安全
	//unsafeRRR();

	// 使用 buf 实现 ioutil.ReadFile 类似效果
	//readCustomAllFileContent()

	// 使用 bufio 实现行统计：
	readLines()

}
