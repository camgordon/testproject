package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	testFunction()

	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "outputoutput")
	file.Close()
}

func testFunction() {
	var x int // Variable declaration. Variables must be declared before use.
	x = 3     // Variable assignment.
	// "Short" declarations use := to infer the type, declare, and assign.
	y := 4
	sum, prod := learnMultiple(x, y)        // Function returns two values.
	fmt.Println("sum:", sum, "prod:", prod) // Simple output.
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // Return two values.
}
