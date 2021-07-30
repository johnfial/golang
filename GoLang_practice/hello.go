// https://golang.org/dl/
// https://www.geeksforgeeks.org/golang/
// https://ispycode.com/GO/Path-and-Filepath/Get-file-name-extension-used-by-path

// tutorials reccomended:
	// https://learnxinyminutes.com/docs/go/
	// and maybe https://gobyexample.com/


// this doc: https://golang.org/doc/tutorial/getting-started

package main
import(
    "fmt"
)
func main() {
    fmt.Println("hello world")
	fmt.Print('this line') 		// .\hello.go:18:12: more than one character in rune literal
	fmt.Print("that line")
	fmt.Printf('number bytes') // .\hello.go:20:13: more than one character in rune literal
}
