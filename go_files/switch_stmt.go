package go_files

import "fmt"

func Switch_stmt() {
	// Reference: https://www.tutorialspoint.com/go/go_switch_statement.htm
	// Expression switch
	var grade string = "B"
	var marks int = 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70 : grade = "C"
		default: grade = "D"  
	}
	switch {
		case grade == "A" :
			fmt.Printf("Excellent!\n" )     
		case grade == "B", grade == "C" :
			fmt.Printf("Well done\n" )      
		case grade == "D" :
			fmt.Printf("You passed\n" )      
		case grade == "F":
			fmt.Printf("Better try again\n" )
		default:
			fmt.Printf("Invalid grade\n" );
	}
	fmt.Printf("Your grade is  %s\n", grade );   



	// type switch
	var x interface{}  // The expression used in a switch statement must have an variable of interface{} type.
	x = 20
     
   switch i := x.(type) {
      case nil:	  
		 fmt.Printf("type of x :%T",i)    
		 // No break is needed in the case statement.            
      case int:	  
         fmt.Printf("x is int\n")                       
      case float64:
         fmt.Printf("x is float64\n")           
      case func(int) float64:
         fmt.Printf("x is func(int)\n")                      
      case bool, string:
         fmt.Printf("x is bool or string\n")       
      default:
         fmt.Printf("don't know the type\n")     
   }  
}

// To run: go run main.go -pkgname=switch_stmt