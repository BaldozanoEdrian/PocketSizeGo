package main

import (
	"bufio"
	"fmt"
	"os"
)

// read file line by line
func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "Error: ")
		os.Exit(1)
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", arg, err)
			continue
		}

		defer f.Close()

		scanner := bufio.NewReader(f)
		for {
			line, err := scanner.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", line, err)
				return
			}
			fmt.Println(line)
		}
	}
}
