package main

import (
	"fmt"

	"github.com/Gopher/basename"
)

func main() {
	src := "src/path/github.com"
	fmt.Println(basename.Basename(src))
}
