// Package called main is used to create executable binary. Program execution starts in package main by calling its function which also called main.
package main

// one import declaration with two ImportSpecs. Each ImportSpec defines single package import.
import (
    "fmt"
	"GoTutorial/stringutil"  // import path is path relative package’s vendor directory or `go env GOPATH`/src
	"GoTutorial/example_package"  // Use actual path while importing and use package specified at the start of the file while using var of func
	// As used my_example.Example_sum
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	fmt.Println(my_example.Example_sum(2, 3))
}
