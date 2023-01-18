package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Conditions/ Cases")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("dicenumber", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Number is 1")
	case 2:
		fmt.Println("Number is 2")
	case 3:
		fmt.Println("Number is 3")
		// to run other too instead of break
		fallthrough
	case 4:
		fmt.Println("Number is 4")
	case 5:
		fmt.Println("Number is 5")
	case 6:
		fmt.Println("Number is 6")
	default:
		fmt.Println("Number is default")
	}

}