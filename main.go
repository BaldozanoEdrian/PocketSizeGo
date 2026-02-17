package main

import "fmt"

const (
	//hasSleep   uint64 = 1 << 0
	hasEaten   uint64 = 1 << 1
	hasStudied uint64 = 1 << 2
	//isHappy    uint64 = 1 << 63
)

func main() {

	currentState := hasStudied | hasEaten

	fmt.Printf("Current State: %d\n", currentState)

	if currentState&hasStudied != 0 {
		fmt.Println("Edrian is happy even if tired!")
	} else {
		fmt.Println("Edrian is disappointed to his self")
	}

}
