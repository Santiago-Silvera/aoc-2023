package aoc2023

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	d []Digit
}

type Digit struct {
	x     int
	y     int
	value int
}

func Day3_part1() {

	nums := make([]Number, 0)
	// for _, num := range nums.n {
	// 	fmt.Printf("%+v\n", num)
	// }

	content, err := os.ReadFile("./inte0st/day3.txt")
	if err != nil {
		println("Error loading the file")
		return
	}

	for y, line := range strings.Split(string(content), "\n") {
		// keep track if we are creating a new number
		newnum := true
		for x, character := range strings.Split(line, "") {
			// Check if character is an int
			if d, err := strconv.Atoi(character); err == nil {
				dig := Digit{x: x, y: y, value: d}
				newnum = false
			} else {
				newnum = true
			}
		}
	}

}
