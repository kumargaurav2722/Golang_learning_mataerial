package main

import (
	"bufio"
	"fmt"
	"os"
)

// standard input is done by using bufio libary and os libarary
// NewReader and Readstrig is used.
// ReadString give the default type of string
// Println contains default "\n" so we have to not give explicitly
// and the below input,_ isystax is called ok comma syntax or error comma syntax
// it is used as a type of try catch
func main() {
	fmt.Println("please enter an rating for Pizza")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)
	fmt.Printf("the type of rating is %T \n", input)
}
