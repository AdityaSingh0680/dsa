package main

import "fmt"

func main() {

	// Define the number of rows for the pyramid
	numRows := 5

	// Outer loop for rows
	for i := 0; i < numRows; i++ {
		// Print spaces for the left side of the pyramid
		for j := 0; j < numRows-i-1; j++ {
			fmt.Print(" ")
		}

		// Print asterisks for the left side of the pyramid
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}

		// Move to the next line after each row
		fmt.Println()
	}
}
