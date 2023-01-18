package main

import (
	"fmt"
)

// we dont have classes in Go lang 
// no inheritance in go lang, no super, no parent 
func main()  {
	fmt.Println("Welcome to structs")
	nikilesh := User{"Nikilesh","nikilesh@gmail.com", true, 30}
	fmt.Printf("%+v\n", nikilesh)
	fmt.Println(nikilesh.Name)

}

// for accessing public Capatalize
type User struct {
	Name string
	Email string
	Status bool
	Age int
}