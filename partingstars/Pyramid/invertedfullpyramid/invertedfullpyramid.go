package main

import "fmt"

func main() {
	rows := 5 // You can change the number of rows as needed

	for i := 0; i < rows; i++ {
		// Print spaces
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}

		// Print asterisks
		for k := 0; k < (rows*2-1)-(2*i); k++ {
			fmt.Print("*")
		}

		// Move to the next line
		fmt.Println()
	}
}
