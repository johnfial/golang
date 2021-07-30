// https://golang.org/dl/
// https://www.geeksforgeeks.org/golang/
// https://ispycode.com/GO/Path-and-Filepath/Get-file-name-extension-used-by-path

// tutorials reccomended:
// https://learnxinyminutes.com/docs/go/
// and maybe https://gobyexample.com/

// this doc: https://golang.org/doc/tutorial/getting-started

// run with:
// "go run hello.go"
// "go run ." // saw this elsewhere, not sure if runs all files, looks for some kind of Go-"project", or what...
package main

import (
	"fmt"
)

var printme = "'variable'" // language looks like JS here, no?

func main() { // main() runs like python's if __name__ == __main__
	fmt.Print("that line")
	fmt.Println("hello world") //prints on SAME line as above, new line AFTER this ends
	fmt.Println(printme)

	// single quotes =(
	// fmt.Print('this line') 		// .\hello.go:18:12: more than one character in rune literal
	// fmt.Printf('number bytes') // .\hello.go:20:13: more than one character in rune literal
}
