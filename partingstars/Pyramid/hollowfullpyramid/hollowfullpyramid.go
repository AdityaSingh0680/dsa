package main

import "fmt"

func main() {
	var rows int

	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {
		// Print spaces for the left side of the pyramid
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// Print asterisks for the left side of the pyramid
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == rows {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		// Move to the next line
		fmt.Println()
	}
}
