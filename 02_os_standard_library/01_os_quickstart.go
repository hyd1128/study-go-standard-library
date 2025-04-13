package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFIle() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

// 创建目录
func createDir() {
	err := os.Mkdir("ms", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 创建多级目录
func mkdirAll() {
	err := os.MkdirAll("ms/one/two", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除一个空的文件
// func removeFile() {
// 	err := os.Remove("test.txt")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }

// 删除一个空的目录
func removeDir() {
	// 删除的是目录路径最底层的目录
	err := os.Remove("ms/one/two")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 强制删除目录以及目录中的文件
func removeAllDir() {
	err := os.RemoveAll("ms")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 获取工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

// 修改工作目录
func chDir() {
	err := os.Chdir("D:\\code_workspace\\self_project\\study-go-standard-library")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	result, _ := os.Getwd()
	fmt.Println(result)
}

// 获取临时目录
func tempDir() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

// 重命名文件
func renameFile() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 重命名目录
func renameDir() {
	err := os.Rename("ms", "ms1")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	// createFIle()
	// createDir()
	// mkdirAll()
	// removeFile()
	// removeDir()
	// removeDir()
	// removeAllDir()
	// getWd()
	// chDir()
	// tempDir()
	renameFile()
	renameDir()
}
