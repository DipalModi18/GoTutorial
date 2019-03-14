package main

import (
	"fmt"
	"strings"
	"os"
	"flag"  // Go provides a flag package supporting basic command-line flag parsing.
)

func main() {
	// Reference: https://gobyexample.com/command-line-arguments
	// os.Args provides access to raw command-line arguments
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	// fmt.Println(os.Args[2])

	// To run: go run os_args.go a b c d



	/* Command-Line Flags */
	// Basic flag declarations are available for string, integer, and boolean options. 
	// Here we declare a string flag word with a default value "foo" and a short description. 
	// This flag.String function returns a string pointer (not a string value)
	wordPtr := flag.String("word", "foo", "a string")  // (key, value, description)
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	
	// It’s also possible to declare an option that uses an existing var declared elsewhere in the program. 
	// Note that we need to pass in a pointer to the flag declaration function.
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	
	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	
	// To run: go run os_args.go -word=opt -numb=7 -fork -svar=flag
	// 		   go run os_args.go -word=opt a1 a2 a3
	// 		   go run os_args.go -h  // for help


	// To set environment variable
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))  // This will return an empty string if the key isn’t present in the environment.

	// os.Environ to list all key/value pairs in the environment.
	// This returns a slice of strings in the form KEY=value.
	// fmt.Println(os.Environ())
	for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Printf("key: %s, value: %s\n", pair[0], pair[1])
    }

}

