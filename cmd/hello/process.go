package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/OneHundred86/hello"
)

func doMain() {
	words := hello.SayHelloTo(name)
	for i := 0; i < count; i++ {
		_, err := os.Stdout.WriteString(words + "\n")

		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
	}
}

func doHelp() {
	flag.Usage()
}

func doVersion() {
	fmt.Printf("Version: %s", VERSION)
}
