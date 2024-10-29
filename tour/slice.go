package tour

import "fmt"

func Slice() {
	// array fixed, slice flexible
	// A slice does not store any data, 
	// it just describes a section of an underlying array.

	// A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	// Can extend a slice's length by re-slicing it, provided it has sufficient capacity.

	// Changing the elements of a slice modifies 
	// the corresponding elements of its underlying array.
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)

	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	var nils []int
	printSlice(nils)
	if nils == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}