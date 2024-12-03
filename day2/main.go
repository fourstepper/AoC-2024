package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func stringSliceToIntSlice(input []string) (output []int) {
	for _, field := range input {
		intField, _ := strconv.Atoi(field)
		output = append(output, intField)
	}
	return output
}

func isGood(fields []int) (sorted, ok bool) {
	reverse := slices.Clone(fields)
	slices.Reverse(reverse)

	if slices.IsSorted(fields) || slices.IsSorted(reverse) {
		sorted = true
	}

	ok = true
	for i := range len(fields) - 1 {
		d := diff(fields[i], fields[i+1])
		if !(d >= 1 && d <= 3) {
			ok = false
		}
	}
	return sorted, ok
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	data := strings.Split(strings.TrimSpace(string(file)), "\n")

	p1 := 0
	p2 := 0
	for _, line := range data {
		input := stringSliceToIntSlice(strings.Fields(line))

		sorted, good := isGood(input)
		if sorted && good {
			p1 += 1
		}

		allGood := false
		for j := range input {
			xs := slices.Clone(input)
			xs = slices.Delete(xs, j, j+1)

			sorted, good := isGood(xs)
			if sorted && good {
				allGood = true
			}
		}
		if allGood {
			p2 += 1
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
