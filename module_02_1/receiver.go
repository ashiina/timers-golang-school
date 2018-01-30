package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

func (c Color) String() string {
	strings := []string{"White", "Black", "Blue", "Red", "Yellow"}
	return strings[c]
}

type Rectangle struct {
	width, height float64
	color         Color
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) SetColor(c Color) {
	r.color = c
}

func main() {
	R := Rectangle{9, 4, WHITE}
	fmt.Println(R.Area())
	R.SetColor(RED)
	fmt.Println(R.color.String())
}
