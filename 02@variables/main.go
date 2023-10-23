package main

import "fmt"

// by using starting  letter capital in var makes it public otherwise it is private
// we cannot use no var style outside of any method or function
var JwtToken = " wkjhgf"

func main() {
	var user string = "kumar"
	fmt.Println("my name is ", user)
	fmt.Printf("my name is of type %T \n", user)

	// type style
	var Myincome = 10000
	fmt.Println("my income is ", Myincome)
	fmt.Printf("my income is of type %T \n", Myincome)

	// without var style

	mySpending := 15000
	fmt.Println("my spending is ", mySpending)
	fmt.Printf("my spending is of type %T", mySpending)

}
