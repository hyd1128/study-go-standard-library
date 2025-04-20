package main

import (
	"fmt"
	"os"
)

// 读操作
// func main() {
// 	f, err := os.Open("a.txt")
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	defer func(f *os.File) {
// 		err := f.Close()
// 		if err != nil {
// 			fmt.Printf("%v", err)
// 		}
// 	}(f)
//
// 	buf := make([]byte, 12)
//
// 	for {
// 		n, err2 := f.Read(buf)
// 		if n == 0 || err2 == io.EOF {
// 			fmt.Println("文件已经读取完毕")
// 			break
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

// 写操作
// func main() {
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0775)
// 	f.Write([]byte("hello golang"))
// 	f.Close()
// }

// Seeker接口
/*
Seeker 用来移动数据的读写指针

Seek 设置下一次读写操作的指针位置，每次的读写操作都是从指针位置开始的

- whence 的含义：

  - 如果 whence 为 0：表示从数据的开头开始移动指针
  - 如果 whence 为 1：表示从数据的当前指针位置开始移动指针
  - 如果 whence 为 2：表示从数据的尾部开始移动指针

- offset 是指针移动的偏移量

  返回移动后的指针位置和移动过程中遇到的任何错误
*/
func main() {
	f, _ := os.Open("a.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
}
