package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	var firstCol []int
	var secondCol []int

	data := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, v := range data {
		row := strings.Fields(v)
		first, _ := strconv.Atoi(row[0])
		second, _ := strconv.Atoi(row[1])

		firstCol = append(firstCol, first)
		secondCol = append(secondCol, second)
	}
	slices.Sort(firstCol)
	slices.Sort(secondCol)

	total := 0
	for idx := range firstCol {
		if firstCol[idx] > secondCol[idx] {
			total += firstCol[idx] - secondCol[idx]
		} else {
			total += secondCol[idx] - firstCol[idx]
		}
	}
	fmt.Println(total)

	similarityScore := 0
	for idx := range firstCol {
		count := 0
		for _, v := range secondCol {
			if firstCol[idx] == v {
				count += 1
			}
		}
		similarityScore += firstCol[idx] * count
	}
	fmt.Println(similarityScore)
}
