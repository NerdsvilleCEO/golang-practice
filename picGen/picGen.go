package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	/*
	  For some reason, if we declare the variables as follows:
	    var x,y int
	  the variables are out of scope when looping
	  TODO: Investigate why
	*/
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		slice := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			val := x ^ 3 + y ^ 2
			slice[x] = uint8(val)
		}
		pic[y] = slice
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
