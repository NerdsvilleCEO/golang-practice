package main

import (
	"fmt"
)

func main() {
	/* Slices are better than arrays for ordered collections */
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors[0], colors[1], colors[2])

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)
	fmt.Println("Number of colors:", len(colors))
}
