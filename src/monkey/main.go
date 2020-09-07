package main

import (
	"fmt"
	"monkey/repl"
)

func main() {
	fmt.Printf("Enter commands to evaluate monkey language\n")
	fmt.Println("-----------------------------------")
	repl.Start()
}
