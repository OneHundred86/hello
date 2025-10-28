package hello

import "fmt"

func SayHelloTo(name string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("hello %s", name)
}
