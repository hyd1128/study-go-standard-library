package main

import (
	"fmt"
	"time"
)

// func checks(s string) (string, error) {
// 	if s == "" {
// 		err := errors.New("字符串不能够为空")
// 		return "", err
// 	} else {
// 		return s, nil
// 	}
// }
//
// func main() {
// 	s, err := checks("")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err.Error())
// 	} else {
// 		fmt.Printf("s: %v\n", s)
// 	}
// }

// 自定义异常
// 自定义error类型
type MyError struct {
	When time.Time // 发生错误的时间
	What string    // 错误文本信息
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
