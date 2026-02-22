package main

import (
	"fmt"

	"github.com/Gopher/basename"
)

func main() {
	src := "src/path/github/myGoProjects.com"
	fmt.Println(basename.Basename(src))
}
