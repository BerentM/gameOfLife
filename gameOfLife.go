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
				b[i][j] = " "
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

func main() {
	univ := NewUniverse()
	univ.Show()
	fmt.Println()
	univ.Seed()
	univ.Show()
}
