package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	fmt.Println(languages)
	fmt.Println(languages["js"])
	delete(languages, "rb")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println(key + " : " + value)
	}
}