package main

import (
		"fmt"
		"tour"
)

func main () {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered panic, ", r)
		}
	}()

	fmt.Println("* * Types * *")
	tour.Types()
	fmt.Println()

	fmt.Println("* * Loops * *")
	tour.ForLoop()
	fmt.Println()

	fmt.Println("* * Conditionals * *")
	tour.IfElse()
	tour.Switch()
	fmt.Println()

	fmt.Println("* * Defer * *")
	tour.DeferEx()
	fmt.Println()

	fmt.Println("* * Panic * *")
	tour.Recover()
	fmt.Println()

	fmt.Println("* * Pointer * *")
	tour.Pointer()
	fmt.Println()

	fmt.Println("* * Struct * *")
	tour.Struct()
	fmt.Println()

	fmt.Println("* * Array * *")
	tour.Array()
	fmt.Println()

	fmt.Println("* * Slice * *")
	tour.Slice()
	fmt.Println()
}