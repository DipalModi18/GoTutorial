// there are two types of packages. 
// An executable package and an utility package. 
// A executable package is your main application since you will be running it. 
// An utility package is not self executable, instead it enhances functionality of an executable package by providing utility functions and other important assets.
// Package called main is used to create executable binary. 
// Program execution starts in package main by calling its function which also called main.
package main

// one import declaration with three ImportSpecs. Each ImportSpec defines single package import.
import (
    "fmt"
	"GoTutorial/stringutil"  // import path is path relative package’s vendor directory or `go env GOPATH`/src
	"GoTutorial/example_package"  // Use actual path while importing and use package specified at the start of the file while using var of func
	// As used my_example.Example_sum
)

var my_global_var int

// init function is called by Go when a package is initialized.
// You can have multiple init functions in a file or a package. 
// Order of the execution of init function in a file will be according to the order of their appearances.
// You can have init function anywhere in the package.
// After all init functions are executed, main function is called. =
// Hence, main job of init function is to initialize global variables that can not be initialized in global context. 
func init() {
	fmt.Println("In First Init of hello.go")
	my_global_var = 10
}

func init() {
	fmt.Println("In Second Init of hello.go")
	my_global_var = 30
}

func main() {
	fmt.Println("In main of hello.go")

	fmt.Println(stringutil.Reverse("!oG ,olleH"))

	fmt.Println(my_example.Example_sum(2, 3))

	fmt.Printf("my_global_var: %d\n", my_global_var)
}


// Reference: https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc