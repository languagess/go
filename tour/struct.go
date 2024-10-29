package tour

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Struct() {
	fmt.Println(Vertex{X: 2, Y: 1})
	v := Vertex{}
	v.X, v.Y = 4, 10
	fmt.Println(v)

	v = Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}