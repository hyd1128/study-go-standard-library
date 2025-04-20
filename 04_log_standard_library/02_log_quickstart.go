package main

import (
	"log"
	"os"
)

// func main() {
// 	i := log.Flags()
// 	fmt.Printf("i: %v\n", i)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
// }

// 输出日志到文件
// func main() {
// 	f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		log.Panic("打开日志文件异常")
// 	}
//
// 	log.SetOutput(f)
// 	log.Print("this is a file log...")
// }

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger = log.New(logFile, "[Mylog] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	logger.Println("自定义logger")
}
