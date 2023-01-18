package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to pointers")

	myNumber := 23
	var ptr = &myNumber
	
	fmt.Println("value pointers", ptr)
	fmt.Println("value pointers", *ptr)
	
}