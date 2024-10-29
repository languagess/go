package tour

import (
	"fmt"
	"runtime"
)

func IfElse() {
	x := 0
	if x < 0 {
		for {
			fmt.Println("Ctrl^C bruh")
		}
	} else {
		fmt.Println("else statement")
	}

	if ok := false; !ok {
		fmt.Println("Not ok")
	}
}

func Switch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin", "whatisthis":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

