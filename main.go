package main

import (
	"fmt"
	"flag"
	"GoTutorial/go_files"
	"GoTutorial/struct_examples"
)

func main() {
	fmt.Println("In main.go")
	pkgname_ptr := flag.String("pkgname", "", "Package to be executed")

	flag.Parse()
	fmt.Println(*pkgname_ptr)

	switch *pkgname_ptr {
	case "loops_example":
		go_files.Loops_example()
	case "structpkg":
		structpkg.Struct_example()
		// go run main.go -pkgname=structpkg
	case "os_args":
		go_files.Os_args()  
		// go run main.go -pkgname=os_args a b c d
	case "switch_stmt":
		go_files.Switch_stmt()
		// go run main.go -pkgname=switch_stmt
	case "variable_declarations":
		go_files.Variable_declarations()
		// go run main.go -pkgname=variable_declarations
	default: 
		fmt.Println("NO files choosen")
	}
}