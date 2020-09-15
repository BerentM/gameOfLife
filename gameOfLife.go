package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Universe [][]bool

const (
	width  int = 80
	height int = 15
)

func NewUniverse() Universe {
	a := make(Universe, height)
	for i := 0; i < len(a); i++ {
		a[i] = make([]bool, width)
	}
	return a
}

func (u Universe) Show() {
	b := make([][]string, height)
	for i := 0; i < len(b); i++ {
		b[i] = make([]string, width)
	}

	for i := 0; i < len(u); i++ {
		for j := 0; j < len(u[i]); j++ {
			if u[i][j] == false {
				b[i][j] = "_"
			} else {
				b[i][j] = "*"
			}
		}
	}
	for _, h := range b {
		fmt.Println(h)
	}
}

func (u Universe) Seed() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(u); i++ {
		for j := 0; j < len(u[i]); j++ {
			if rand.Intn(100) < 25 {
				u[i][j] = true
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	if x < 0 {
		x = width - 1
	} else if x > width-1 {
		x = 0
	}

	if y < 0 {
		y = height - 1
	} else if y > height-1 {
		y = 0
	}

	if u[y][x] == true {
		// fmt.Printf("x=%v, max(x)=%v, y=%v, max(y)=%v\ncell status: alive\n", x, width, y, height)
		return true
	} else {
		// fmt.Printf("x=%v, max(x)=%v, y=%v, max(y)=%v\ncell status: dead\n", x, width, y, height)
		return false
	}
}

func (u Universe) Neighbours(x, y int) int {
	var cnt int
	x--
	y--
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (i != 1) || (j != 1) {
				if u.Alive(x, y) == true {
					cnt++
				}
			}
			y++
		}
		x++
		y -= 3
	}
	// fmt.Printf("Liczba żywych: %v\n", cnt)
	return cnt
}

func (u Universe) Next(x, y int) bool {
	var out bool
	if (u.Alive(x, y) == true) && u.Neighbours(x, y) < 2 {
		out = false
	} else if (u.Alive(x, y) == true) && (u.Neighbours(x, y) >= 2) && (u.Neighbours(x, y) < 4) {
		out = true
	} else if (u.Alive(x, y) == true) && u.Neighbours(x, y) > 3 {
		out = false
	} else if (u.Alive(x, y) == false) && (u.Neighbours(x, y) == 3) {
		out = true
	}
	return out
}

func Step(a, b Universe) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(i, j)
		}
	}
	a, b = b, a
}

func main() {
	univ := NewUniverse()
	univb := NewUniverse()
	univ.Seed()
	for i := 0; i < 10; i++ {
		fmt.Print("\x0c")
		time.Sleep(3 * time.Second)
		univ.Show()
		Step(univ, univb)
	}
}
