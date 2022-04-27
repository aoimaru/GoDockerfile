package main

import (
	"fmt"

	lib "github.com/aoimaru/GoDockerfile/lib"
)

func main() {
	fmt.Println("Hello World")
	out := lib.Hogehoge("nakamura")
	fmt.Println(out)
}
