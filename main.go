package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "stop" || text == "exit" {
			break
		}
		counts[text]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		} else {
			fmt.Printf("No Duplicate: %s\t%d\n", line, n)
		}
	}
}
