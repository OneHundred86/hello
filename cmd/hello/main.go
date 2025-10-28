package main

import (
	"flag"
	"os"

	"github.com/OneHundred86/hello"
)

var (
	name  string
	count int
)

func init() {
	// 定义命令行参数
	flag.StringVar(&name, "name", "默认值", "这是使用说明")
	flag.IntVar(&count, "cnt", 1, "说话次数")
}

func main() {
	// 解析命令行参数
	flag.Parse()

	words := hello.SayHelloTo(name)
	for i := 0; i < count; i++ {
		_, err := os.Stdout.WriteString(words + "\n")

		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
	}
}
