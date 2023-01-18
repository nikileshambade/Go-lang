package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to slices")

	var courses = []string{"reactjs", "solidjs", "vuejs", "angular"}

	fmt.Println(courses)
	// append can be used to add as well as remove values
	// last value is non inclusive
	var index int = 3
	courses = append(courses[:index])
	fmt.Println(courses)
	index = 1
	courses = append(courses[:index], courses[index + 1:]...)
	fmt.Println(courses)

}