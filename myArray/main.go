package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to Array")

	var list [4]string
	list[0] = "Nikilesh"
	list[1] = "Ambade"

	fmt.Println("List : ", list);

	// slice

	var sliceObj = []string{ "apple", "tomato" }
	fmt.Println("Slice : ", sliceObj);
	fmt.Println("Slice : ", len(sliceObj));

	sliceObj = append(sliceObj, "newone", "anotherone")
	fmt.Println("Slice : ", sliceObj);
	// splice array start: destination [destination is non inclusive]
	fmt.Println("Slice : ", append(sliceObj[1:]))
	fmt.Println("Slice : ", append(sliceObj[1:2]))

	// dynamic memory allocation
	highScore := make([]int, 4)

	highScore[0] = 233
	highScore[1] = 234
	highScore[2] = 563
	highScore[3] = 64

	fmt.Println(highScore)

	highScore = append(highScore, 435, 434)
	fmt.Println(highScore)
}