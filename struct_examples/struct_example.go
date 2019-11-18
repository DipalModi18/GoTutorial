// Reference: https://gobyexample.com/structs
// Go structs are useful for grouping data together to form records.
package structpkg

import "fmt"

type Person struct {  // This Person struct type has name and age fields.
	name string
	age int
}

func Struct_example() {
	fmt.Println(Person{name:"ddd", age:20})
	fmt.Println(Person{name:"ddd"})  // Omitted fields will be zero-valued
	fmt.Println(Person{"Bob", 20})
	fmt.Println(&Person{name: "Ann", age: 40})  // An & prefix yields a pointer to the struct.

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)
	
	sp := &s  // Pointer to the struct object
	fmt.Println(sp.age)
	
	sp.age = 51  // Structs are mutable.
	fmt.Println(sp.age)
	

	rectangle1 := Rect{width: 4, height:5}  // From struct_method.go
	fmt.Printf("Area of the rectangle: %d\n", rectangle1.area())
	fmt.Printf("Perimeter of the rectangle: %d\n", rectangle1.perim())

	rp := &rectangle1
    fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	
	Interface_example()
}

// To run: go run main.go -pkgname=structpkg