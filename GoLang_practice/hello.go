// https://golang.org/dl/
// https://www.geeksforgeeks.org/golang/
// https://ispycode.com/GO/Path-and-Filepath/Get-file-name-extension-used-by-path

// tutorials reccomended:
// https://learnxinyminutes.com/docs/go/
// and maybe https://gobyexample.com/

// this doc: https://golang.org/doc/tutorial/getting-started

// / / / / /  / / / / /  / / / / /  / / / / /  / / / / /  / / / / /
// SETUP
// (install go, obviously)
// to run: BASH > cd > "go run hello.go"
// "go run ." // saw this elsewhere, not sure if runs all files, looks for some kind of Go-"project", or what...
// ^^ does not run the hello.go file here if ran in same folder.
// / / / / /  / / / / /  / / / / /  / / / / /  / / / / /  / / / / /
// don't want to run this here? Run on the web @  https://play.golang.org/
// / / / / /  / / / / /  / / / / /  / / / / /  / / / / /  / / / / /

package main

import (
	"fmt"
	m "math"
	"os"
	// very annoing: cntrl + S in VS Code removes import lines with syntax problems...
)

var printme = "'variable'"                         // language looks like JS here, no?
var concatenate_test = "testing" + "555" + printme // NOTE needed "" for string

func main() { // main() runs like python's if __name__ == __main__
	fmt.Print("that line")
	fmt.Println("hello world") //prints on SAME line as above, new line AFTER this ends
	fmt.Println(printme)
	fmt.Println(concatenate_test)

	fmt.Print(fmt.Print()) // prints None output, as print returned nothing: 0 <nil>
	fmt.Print(fmt.Print)   // prints function location: 0xfa1540, for example

	// single quotes =(
	// fmt.Print('this line') 		// .\hello.go:18:12: more than one character in rune literal
	// fmt.Printf('number bytes') // .\hello.go:20:13: more than one character in rune literal
}

func beyondHello() {
	var x int
	x = 33
	y := 4 // wtf does := do? // TODO
	sum, prod := learnMultiple(x, y) // returns two values
	fmt.Println("sum is", sum, "product is:" prod) // looks like normal concatenation
	// learnTypes() // wtf does this do?
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}