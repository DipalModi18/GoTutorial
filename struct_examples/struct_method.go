// Reference: https://gobyexample.com/methods
package structpkg

type Rect struct {
	height, width int
}

// This area method has a receiver type of *Rect.
func (r *Rect) area() int{
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
func (r Rect) perim() int {
    return 2*r.width + 2*r.height
}