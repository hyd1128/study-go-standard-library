package main

// func main() {
// 	var b = []byte("seafood")
//
// 	a := bytes.ToUpper(b)
// 	fmt.Println(a, b)
//
// 	c := b[0:4]
// 	c[0] = 'A'
// 	fmt.Println(c, b)
// }

// func main() {
// 	s1 := "Φφϕ kKK"
// 	s2 := "ϕΦφ KkK"
//
// 	// 看看 s1 里面是什么
// 	for _, c := range s1 {
// 		fmt.Printf("%-5x", c)
// 	}
// 	fmt.Println()
// 	// 看看 s2 里面是什么
// 	for _, c := range s2 {
// 		fmt.Printf("%-5x", c)
// 	}
// 	fmt.Println()
// 	// 看看 s1 和 s2 是否相似
// 	fmt.Println(bytes.EqualFold([]byte(s1), []byte(s2)))
// }

// func main() {
// 	bs := [][]byte{ // [][]byte 字节切片 二维数组
// 		[]byte("Hello World !"),
// 		[]byte("Hello 世界！"),
// 		[]byte("hello golang ."),
// 	}
//
// 	f := func(r rune) bool {
// 		return bytes.ContainsRune([]byte("!！. "), r) // 判断r字符是否包含在"!！. "内
// 	}
//
// 	for _, b := range bs { // range bs  取得下标和[]byte
// 		fmt.Printf("去掉两边: %q\n", bytes.TrimFunc(b, f)) // 去掉两边满足函数的字符
// 	}
//
// 	for _, b := range bs {
// 		fmt.Printf("去掉前缀: %q\n", bytes.TrimPrefix(b, []byte("Hello "))) // 去掉前缀
// 	}
// }

// func main() {
// 	// 定义一个字节数组
// 	b1 := []byte("hello")
// 	var b2 []byte = []byte("hello") // 显示的转换
// 	fmt.Printf("%v", string(b1))
// 	fmt.Printf("%v", string(b2))
//
// }

// func main() {
// 	b := []byte("hello golang") // 字符串强转为byte切片
// 	sublice1 := []byte("hello")
// 	sublice2 := []byte("Hello")
// 	fmt.Println(bytes.Contains(b, sublice1)) // true
// 	fmt.Println(bytes.Contains(b, sublice2)) // false
//
// 	s := []byte("hellooooooooo")
// 	sep1 := []byte("h")
// 	sep2 := []byte("l")
// 	sep3 := []byte("o")
// 	fmt.Println(bytes.Count(s, sep1)) // 1
// 	fmt.Println(bytes.Count(s, sep2)) // 2
// 	fmt.Println(bytes.Count(s, sep3)) // 9
// }

// buffer
// func main() {
// 	rd := bytes.NewBufferString("Hello World!")
// 	buf := make([]byte, 6)
// 	// 获取数据切片
// 	b := rd.Bytes()
// 	// 读出一部分数据，看看切片有没有变化
// 	rd.Read(buf)
// 	fmt.Printf("%s\n", rd.String())
// 	fmt.Printf("%s\n\n", b)
//
// 	// 写入一部分数据，看看切片有没有变化
// 	rd.Write([]byte("abcdefg"))
// 	fmt.Printf("%s\n", rd.String())
// 	fmt.Printf("%s\n\n", b)
//
// 	// 再读出一部分数据，看看切片有没有变化
// 	rd.Read(buf)
// 	fmt.Printf("%s\n", rd.String())
// 	fmt.Printf("%s\n", b)
// }

// func main() {
// 	data := "123456789"
// 	// 通过[]byte创建Reader
// 	re := bytes.NewReader([]byte(data))
// 	// 返回未读取部分的长度
// 	fmt.Println("re len : ", re.Len())
// 	// 返回底层数据总长度
// 	fmt.Println("re size : ", re.Size())
//
// 	fmt.Println("---------------")
//
// 	buf := make([]byte, 2)
// 	for {
// 		// 读取数据
// 		n, err := re.Read(buf)
// 		if err != nil {
// 			fmt.Printf("%v\n", err)
// 			break
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

// func main() {
// 	data := "123456789"
// 	// 通过[]byte创建Reader
// 	re := bytes.NewReader([]byte(data))
//
// 	buf := make([]byte, 2)
//
// 	re.Seek(1, 0)
// 	// 设置偏移量
// 	for {
// 		// 一个字节一个字节的读
// 		b, err := re.ReadByte()
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(b))
// 	}
// 	fmt.Println("----------------")
//
// 	re.Seek(0, 0)
// 	off := int64(0)
// 	for {
// 		// 指定偏移量读取
// 		n, err := re.ReadAt(buf, off)
// 		if err != nil {
// 			break
// 		}
// 		off += int64(n)
// 		fmt.Println(off, string(buf[:n]))
// 	}
// }
