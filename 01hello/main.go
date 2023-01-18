package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

const NumberOfUser int = 300
func main()  {
	numberOfUser := "string"
	fmt.Println(numberOfUser)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("read a string")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading, ", input)
	fmt.Printf("Thanks for rating %T", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64);
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to input ", numRating + 1)
	}
}