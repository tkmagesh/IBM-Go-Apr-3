//package declaration
package main

//package imports
import "fmt"

//package variables

//package init fn

//function main
func main() {
	/*
		var message string
		message = "Hello World!"
	*/

	/*
		var message string = "Hello World!"
	*/

	message := "Hello World!"
	fmt.Println(message)

	/*
		var userName string
		var greetMsg string
		userName = "Magesh"
		greetMsg = "Have a nice day!"
	*/

	var (
		userName string
		greetMsg string
	)
	userName = "Magesh"
	greetMsg = "Have a nice day!"

	/*
		var userName, greetMsg string
		userName = "Magesh"
		greetMsg = "Have a nice day!"
	*/
	/*
		var userName, greetMsg string = "Magesh", "Have a nice day!"
	*/
	/*
		userName, greetMsg := "Magesh", "Have a nice day!"
	*/
	fmt.Println(userName, greetMsg)
}

//other types and functions
