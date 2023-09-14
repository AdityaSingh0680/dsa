package main

import "fmt"

func main() {
	rows := 5 // Change the number of rows as needed

	for i := 1; i <= rows; i++ {
		// Print spaces before the first asterisk
		for j := 1; j < i; j++ {
			fmt.Print(" ")
		}

		// Print the first asterisk for each row
		fmt.Print("*")

		// Print spaces between the first and last asterisk
		for j := 1; j <= (2 * (rows - i - 1)); j++ {
			fmt.Print(" ")
		}

		// Print the last asterisk for each row
		if i != 1 {
			fmt.Print("*")
		}

		// Move to the next line
		fmt.Println()
	}
}
