package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	// f, err := os.Open("a.txt")

	// NopCloser
	// if err != nil {
	// 	fmt.Printf("err: %v", err)
	// }
	// readCloser := ioutil.NopCloser(f)
	// fmt.Printf("readCloser: %v\n", readCloser)

	// readAll
	// b, err := ioutil.ReadAll(f)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	// fmt.Printf("string(b): %v\n", string(b))

	// readDir
	// f1, _ := ioutil.ReadDir(".")
	// for _, v := range f1 {
	// 	fmt.Printf("v.Namae: %v\n", v.Name())
	// }

	// 直接读取某一个文件 不需要os.Open() 这个文件指针
	// b, _ := ioutil.ReadFile("a.txt")
	// fmt.Printf("string(b): %v\n", string(b))

	// WriteFile
	// ioutil.WriteFile("b.txt", []byte("hello world"), 0664)

	// TempDir
	// content := []byte("temporary file's content")
	// dir, err := ioutil.TempDir(".", "example")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Printf("dir: %v\n", dir)
	// // defer os.RemoveAll(dir) // 销毁临时目录
	//
	// tmpfn := filepath.Join(dir, "tmpfile")
	// if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
	// 	log.Fatal(err)
	// }

	// TempFile()
	// content := []byte("temporary file's content")
	// tmpfile, err := ioutil.TempFile(".", "example")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("tmpfile.Name(): %v\n", tmpfile.Name())
	//
	// // defer os.Remove(tmpfile.Name()) // 销毁临时文件
	//
	// if _, err := tmpfile.Write(content); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := tmpfile.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// s := strings.NewReader("ABCDEFG")
	// str := strings.NewReader("12345")
	// br := bufio.NewReader(s)
	// b, _ := br.ReadString('\n')
	// fmt.Println(b)
	// br.Reset(str)
	// b, _ = br.ReadString('\n')
	// fmt.Println(b)

	// s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	// br := bufio.NewReader(s)
	// p := make([]byte, 10)
	//
	// for {
	// 	n, err := br.Read(p)
	// 	if err == io.EOF {
	// 		break
	// 	} else {
	// 		fmt.Printf("string(p[0:n]): %v\n", string(p[0:n]))
	// 	}
	// }

	// b := bytes.NewBuffer(make([]byte, 0))
	// bw := bufio.NewWriter(b)
	// bw.WriteString("123456789")
	// c := bytes.NewBuffer(make([]byte, 0))
	// bw.Reset(c)
	// bw.WriteString("456")
	// bw.Flush()
	// fmt.Println(b)
	// fmt.Println(c)

	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) // 4096
	fmt.Println(bw.Buffered())  // 0

	bw.WriteString("ABCDEFGHIJKLMN")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)
}
