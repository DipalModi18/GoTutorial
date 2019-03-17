// Reference: https://www.tutorialspoint.com/go/go_variables.htm
package go_files

import "fmt"

func Variable_declarations() {
   var x float64  // A static type variable declaration provides assurance to the compiler that there is one variable available with the given type and name so that the compiler can proceed for further compilation without requiring the complete detail of the variable.
   x = 20.0
   // Or above two line can be combined as: var x float64 = 20.0

   // A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it.
   y := 42 

   // Notice, in case of type inference, we initialized the variable y with := operator, whereas x is initialized using = operator.
   fmt.Println(x)
   fmt.Println(y)
   fmt.Printf("x is of type %T\n", x)
   fmt.Printf("y is of type %T\n", y)	





   // Variables of different types can be declared in one go using type inference.
   var a, b, c = 3, 4, "foo"  
	
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)


   /* There are two kinds of expressions in Go −

      lvalue − Expressions that refer to a memory location is called "lvalue" expression. An lvalue may appear as either the left-hand or right-hand side of an assignment.

      rvalue − The term rvalue refers to a data value that is stored at some address in memory. An rvalue is an expression that cannot have a value assigned to it which means an rvalue may appear on the right- but not left-hand side of an assignment.

   Variables are lvalues and so may appear on the left-hand side of an assignment. Numeric literals are rvalues and so may not be assigned and can not appear on the left-hand side
   */
}

// To run: go run main.go -pkgname=variable_declarations