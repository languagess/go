package tour

import "fmt"

func runAtTheEnd(x int) {
	fmt.Println("I run at the end ", x)
}

func DeferEx() {
	// LIFO
	x := 5
	for x > 0 {
		defer runAtTheEnd(x)
		x -= 1
	}
	fmt.Println("Deferred a func")

	i := 0
	// evaluated here
    defer fmt.Println(i)
    i++
    return
}