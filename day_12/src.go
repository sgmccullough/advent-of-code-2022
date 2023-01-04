package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	part1()

	// part2()
}

func part1() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var elevationMap [][]int

	pos_x := 0
	pos_y := 0

	goal_x := 0
	goal_y := 0

	i := 0
	for scanner.Scan() {
		line := []int{}
		for j, r := range scanner.Text() {
			if r == 83 {
				pos_x = int(j)
				pos_y = int(i)
				r = 115
			} else if r == 69 {
				goal_x = int(j)
				goal_y = int(i)
				r = 101
			}
			line = append(line, int(r-97))
		}
		elevationMap = append(elevationMap, line)
		i++
	}

	fmt.Println(pos_x, pos_y, goal_x, goal_y)
	fmt.Println(elevationMap)

	count := 0
	for {
		if pos_x == goal_x && pos_y == goal_y {
			break
		}
		count++
	}

	fmt.Printf("Part 1: %v\n", count)
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	x := 1 // also the cursor value
	clock := 1

	var screen [240]int
	for i := 0; i < len(screen); i++ {
		screen[i] = 0
	}

	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		curr := strings.Split(line, " ")
		command := curr[0]
		if command == "addx" {
			number, _ := strconv.Atoi(curr[1])
			if clock%40 == 0 {
				lineNumber++
			}
			clock++
			x += number
		}
		if clock%40 == 0 {
			lineNumber++
		}
		clock++
	}

	fmt.Println("Part 2:")
	for i := 0; i < len(screen); i++ {
		if screen[i] == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if i%40 == 39 && i != 0 {
			fmt.Println()
		}
	}
}
