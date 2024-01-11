package aoc2023

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)
const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func Day2_part1() {
	content, err := os.ReadFile("./input/d2.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	sum := 0

	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}

		pregame := strings.Split(line, ":")
		game_id, err := strconv.Atoi(strings.Split(pregame[0], " ")[1])
		if err != nil {
			continue
		}
		validGame := true
		println(game_id)

		for _, rounds := range strings.Split(pregame[1], ";") {
			println("--", rounds)

			for _, cube := range strings.Split(rounds, ",") {

				cubeAmount, err := strconv.Atoi(strings.Split(cube, " ")[1])
				if err != nil {
					println("Error of int conversion")
					continue
				}

				cubeType := strings.Split(cube, " ")[2]
				println("----cubeAmount:", cubeAmount)
				println("----cubeType:", cubeType)
				switch cubeType {
				case "red":
					if cubeAmount > MAX_RED {
						println("Max exceeded")
						validGame = false
					}
				case "green":
					if cubeAmount > MAX_GREEN {
						println("Max exceeded")
						validGame = false
					}
				case "blue":
					if cubeAmount > MAX_BLUE {
						println("Max exceeded")
						validGame = false
					}
				default:
					println("weird cubeType")
				}
			}
		}
		if validGame {
			sum += game_id
		}
	}
	println(sum)
}

func Day2_part2() {
	content, err := os.ReadFile("./input/d2.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	sum := 0

	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}

		pregame := strings.Split(line, ":")

		min_red := 0
		min_green := 0
		min_blue := 0
		for _, rounds := range strings.Split(pregame[1], ";") {

			for _, cube := range strings.Split(rounds, ",") {

				cubeAmount, err := strconv.Atoi(strings.Split(cube, " ")[1])
				if err != nil {
					println("Error of int conversion")
					continue
				}

				cubeType := strings.Split(cube, " ")[2]
				switch cubeType{
				case "red":
					if cubeAmount > min_red {
						min_red = cubeAmount
					}
				case "green":
					if cubeAmount > min_green {
						min_green = cubeAmount
					}
				case "blue":
					if cubeAmount > min_blue {
						min_blue = cubeAmount
					}
				}
			}
		}
		power := min_blue * min_green * min_red
		sum += power
	}
	println(sum)
}

