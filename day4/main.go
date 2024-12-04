package main

import (
	"fmt"
	"os"
	"strings"
)

func getData(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	return strings.Split(strings.TrimSpace(string(file)), "\n")
}
func xmasInc(m, a, s string) bool {
	xmas := m + a + s
	if xmas == "MAS" {
		return true
	}
	return false
}

type Pos struct {
	X int
	Y int
}

func main() {
	data := getData("input.txt")
	mapped := map[Pos]string{}

	for y, row := range data {
		for x, col := range row {
			mapped[Pos{X: x, Y: y}] = string(col)
		}
	}

	p1 := 0
	for k, v := range mapped {
		if v == "X" {
			// Straights
			if xmasInc(
				mapped[Pos{X: k.X + 1, Y: k.Y}],
				mapped[Pos{X: k.X + 2, Y: k.Y}],
				mapped[Pos{X: k.X + 3, Y: k.Y}],
			) {
				p1 += 1
			}
			if xmasInc(
				mapped[Pos{X: k.X - 1, Y: k.Y}],
				mapped[Pos{X: k.X - 2, Y: k.Y}],
				mapped[Pos{X: k.X - 3, Y: k.Y}],
			) {
				p1 += 1
			}
			if xmasInc(
				mapped[Pos{X: k.X, Y: k.Y - 1}],
				mapped[Pos{X: k.X, Y: k.Y - 2}],
				mapped[Pos{X: k.X, Y: k.Y - 3}],
			) {
				p1 += 1
			}
			if xmasInc(
				mapped[Pos{X: k.X, Y: k.Y + 1}],
				mapped[Pos{X: k.X, Y: k.Y + 2}],
				mapped[Pos{X: k.X, Y: k.Y + 3}],
			) {
				p1 += 1
			}

			// Diagonals
			if xmasInc(
				mapped[Pos{X: k.X + 1, Y: k.Y + 1}],
				mapped[Pos{X: k.X + 2, Y: k.Y + 2}],
				mapped[Pos{X: k.X + 3, Y: k.Y + 3}],
			) {
				p1 += 1
			}
			if xmasInc(
				mapped[Pos{X: k.X - 1, Y: k.Y - 1}],
				mapped[Pos{X: k.X - 2, Y: k.Y - 2}],
				mapped[Pos{X: k.X - 3, Y: k.Y - 3}],
			) {
				p1 += 1
			}
			if xmasInc(
				mapped[Pos{X: k.X + 1, Y: k.Y - 1}],
				mapped[Pos{X: k.X + 2, Y: k.Y - 2}],
				mapped[Pos{X: k.X + 3, Y: k.Y - 3}],
			) {
				p1 += 1
			}
			if xmasInc(
				mapped[Pos{X: k.X - 1, Y: k.Y + 1}],
				mapped[Pos{X: k.X - 2, Y: k.Y + 2}],
				mapped[Pos{X: k.X - 3, Y: k.Y + 3}],
			) {
				p1 += 1
			}
		}
	}
	fmt.Println(p1)
}
