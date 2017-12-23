package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    p :=  make([][]uint8, dy)
	for y := 0; y < dy; y++ {
	    var px []uint8
		for x := 0; x < dx; x++ {
			px = append(px, uint8((x^y)))
		}		
		p[y] = px
	}
	return p
}

func main() {
	pic.Show(Pic)
}
