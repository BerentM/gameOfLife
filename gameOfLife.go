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
	x = (x + width) % width
	y = (y + height) % height

	return u[y][x]
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

	return cnt
}

func (u Universe) Next(x, y int) bool {
	alive := u.Alive(x, y)
	neighbours := u.Neighbours(x, y)
	if alive && (neighbours < 2) {
		return false
	} else if alive && (neighbours > 3) {
		return false
	} else if neighbours == 3 {
		return true
	}

	return alive
}

func Step(a, b Universe) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(i, j)
		}
	}
}

func main() {
	univ := NewUniverse()
	univb := NewUniverse()
	univ.Seed()
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		univ.Show()
		Step(univ, univb)
		univ, univb = univb, univ
	}
}
