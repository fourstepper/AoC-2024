package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getData(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	return strings.TrimSpace(string(file))
}

func findMulsAndProcess(line string) (total int) {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(line, -1)
	for _, match := range matches {
		r := regexp.MustCompile(`\d+`)
		matches := r.FindAllString(match, -1)
		d0, _ := strconv.Atoi(matches[0])
		d1, _ := strconv.Atoi(matches[1])
		total += d0 * d1
	}
	return total
}

func main() {
	data := getData("input.txt")

	p1 := 0
	p2 := 0

	p1 += findMulsAndProcess(data)

	var enabled_instructions []string
	dos := regexp.MustCompile(`do\(\)`)
	splitByDos := dos.Split(data, -1)

	donts := regexp.MustCompile(`don't\(\).*`)
	for _, v := range splitByDos {
		enabled_instructions = append(enabled_instructions, donts.Split(v, -1)[0])
	}

	for _, v := range enabled_instructions {
		p2 += findMulsAndProcess(v)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
