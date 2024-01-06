package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// day1_part1()
	day1_part2()
}
func day1_part1() {
	content, err := ioutil.ReadFile("../input/d1p1.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	sum := 0
	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}
		var l []int
		for _, c := range strings.Split(line, "") {
			if i, err := strconv.Atoi(c); err == nil {
				l = append(l, i)
			}
		}
		sum += l[0]*10 + l[len(l)-1]
	}
	fmt.Print(sum)
}

func day1_part2() {
	content, err := ioutil.ReadFile("../input/d1p1.txt")
	if err != nil {
		fmt.Print(err)
		return
	}

	sum := 0
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}

		// fmt.Println("This is line: ", linenum)
		fmt.Println(line)
		res := [2]int{0, 0}

		for j := 0; j < 10; j++ {
			line = strings.Replace(line, fmt.Sprint(j), m[j]+fmt.Sprint(j)+m[j], -1)
		}
		minidx := len(line)
		maxidx := 0
		for idx := 1; idx < 10; idx++ {
			// println(m[idx], ":", strings.Index(line, m[idx]))
			// println("minidx:", minidx)
			// println("maxidx:", maxidx)
			if strings.Index(line, m[idx]) < minidx && strings.Index(line, m[idx]) >= 0 {
				minidx = strings.Index(line, m[idx])
				res[0] = idx
			}
			if strings.LastIndex(line, m[idx]) > maxidx && strings.LastIndex(line, m[idx]) >= 0 {
				maxidx = strings.LastIndex(line, m[idx])
				res[1] = idx
			}
		}
		println(res[0], res[1])
		// println(line)
		sum += res[0]*10 + res[1]
	}
	fmt.Print(sum)
}
