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

	part2()
}

func part1() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	x := 1
	clock := 1
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		curr := strings.Split(line, " ")
		command := curr[0]
		if command == "addx" {
			number, _ := strconv.Atoi(curr[1])
			clock++
			sum += check(clock, x)
			x += number
		}

		clock++
		sum += check(clock, x)
	}

	fmt.Printf("Part 1: %v\n", sum)
}

func check(clock int, x int) int {
	if clock == 20 {
		return clock * x
	} else if (clock-20)%40 == 0 {
		return clock * x
	}
	return 0
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
			screen = cycle(clock, x, lineNumber, screen)
			clock++
			x += number
		}
		if clock%40 == 0 {
			lineNumber++
		}
		screen = cycle(clock, x, lineNumber, screen)
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

func cycle(clock int, x int, lineNumber int, screen [240]int) [240]int {
	lineAdustment := x + (40 * lineNumber)
	if lineAdustment-1 > 0 && lineAdustment-1 < 240 && lineAdustment-1 == clock {
		screen[lineAdustment-1] = 1
	} else if lineAdustment > 0 && lineAdustment < 240 && lineAdustment == clock {
		screen[lineAdustment] = 1
	} else if lineAdustment+1 > 0 && lineAdustment+1 < 240 && lineAdustment+1 == clock {
		screen[lineAdustment+1] = 1
	}
	return screen
}
