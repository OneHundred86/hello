package main

import (
	"flag"
)

var (
	name  string
	count int
)

func init() {
	// 定义命令行参数
	flag.StringVar(&name, "name", "默认值", "这是使用说明")
	flag.IntVar(&count, "cnt", 1, "说话次数")

	// 解析命令行参数
	flag.Parse()
}

func main() {
	// fmt.Printf("args: %#v, %#v \n", flag.Args(), flag.Arg(0))
	childCmd := flag.Arg(0)

	switch childCmd {
	case "help":
		doHelp()
	case "version":
		doVersion()
	case "":
		doMain()
	default:
		doHelp()
	}
}
