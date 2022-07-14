package main

import (
	"fmt"
	"os"
)

//global var: (always with '=' never with :=)
var global_var = 1

//types of vars:  (var + name + type = ...)
var number int = 8
var unumber uint = 20
var floatn float32 = 3.14
var isItRight bool = false
var name string = "Billy"

//You don't need to write type after a name of variable

//const vars
const Pi float32 = 3.14

func main() {

	aboba := "local var " //short init for local variables
	fmt.Println("aboba " + aboba)
	fmt.Printf("Pi number is %f \n", Pi)
	fmt.Println("check ", int(unumber)%number)
	fmt.Scan("fdfdfd", aboba)

	n, err := fmt.Fprintln(os.Stdout, name, "is", 13, "years old.")

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

}
