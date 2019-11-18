// Reference: https://gobyexample.com/interfaces
package structpkg

import (
	"fmt"
)

type shape interface { // basic interface for geometric shapes.
	area() float32
	perim() float32
}

type rect struct {
	width, height float32
}

// To implement an interface in Go, we just need to implement all the methods in the interface.
func (r rect) area() float32 {
	return r.width * r.height
}

func (r rect) perim() float32 {
	return 2*r.width + 2*r.height
}

// If a variable has an interface type, then we can call methods that are in the named interface.
func measure(g shape) {
	fmt.Println("geometry: ", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perimeter: ", g.perim())
}

func Interface_example() {
	r := rect{width: 3, height: 4}
	measure(r)

	var s shape
	s = rect{width: 5, height: 6}
	measure(s)
}
