package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Hello")
	currentTime := time.Now()

	fmt.Println("Current time ", currentTime)
	fmt.Println("Current time ", currentTime.Format("2006"))
}