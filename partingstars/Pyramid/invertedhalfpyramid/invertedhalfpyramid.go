package main

import "fmt"

func main() {
	rows := 5 // You can change the number of rows as per your requirement

	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
