package solidrectangle

import "fmt"

func main() {
	width := 5
	height := 10

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
