package my_example

// Go exports a variable if a variable name starts with Uppercase. 
// All other variables not starting with an uppercase letter is private to the package.
func Example_sum(a int, b int) int{
	return a +b
}

// A package name is name of the directory contained in src directory. 
// Hence, go install GoTutorial command looked for GoTutorial sub-directory inside src directory of GOPATH. 
// Then it compiled the package and created GoTutorial binary executable file inside bin directory 
	// which should be executable from terminal since bin directory in the PATH.

// Package declaration which should be first line of code like package main in above example, 
	// can be different than package name. 
// Hence, you might find some packages where package name (name of the directory) is different than package declaration. 
// When you import a package, package declaration is used to create package reference variable

// go install <package> command looks for any file with main package declaration inside given package directory.
// If it finds a file, then Go knows this is an executable program and it needs to create a binary file. 
// A package can have many files but only one file with main function, since that file will be entry point of the execution.

// If a package does not contain file with main package declaration, then Go creates a package archive (.a) file inside pkg directory.