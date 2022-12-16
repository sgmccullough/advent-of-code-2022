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

type Position struct {
	X int
	Y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	head := Position{
		X: 0,
		Y: 0,
	}
	tail := Position{
		X: 0,
		Y: 0,
	}

	positions := []Position{}

	positions = append(positions, Position{X: 0, Y: 0})

	xMax := 0
	yMax := 0

	xMin := 0
	yMin := 0

	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		direction := move[0]
		spaces, _ := strconv.Atoi(move[1])
		for i := 0; i < spaces; i++ {
			if direction == "U" {
				head.Y++
				if Abs(tail.Y-head.Y) > 1 && Abs(tail.X-head.X) > 0 {
					tail.X = head.X
				}
				if Abs(tail.Y-head.Y) > 1 {
					tail.Y++
				}
			} else if direction == "D" {
				head.Y--
				if Abs(tail.Y-head.Y) > 1 && Abs(tail.X-head.X) > 0 {
					tail.X = head.X
				}
				if Abs(tail.Y-head.Y) > 1 {
					tail.Y--
				}
			} else if direction == "L" {
				head.X--
				if Abs(tail.X-head.X) > 1 && Abs(tail.Y-head.Y) > 0 {
					tail.Y = head.Y
				}
				if Abs(tail.X-head.X) > 1 {
					tail.X--
				}
			} else if direction == "R" {
				head.X++
				if Abs(tail.X-head.X) > 1 && Abs(tail.Y-head.Y) > 0 {
					tail.Y = head.Y
				}
				if Abs(tail.X-head.X) > 1 {
					tail.X++
				}
			}
			positions = append(positions, Position{X: tail.X, Y: tail.Y})
			if tail.X > xMax {
				xMax = tail.X
			}
			if tail.Y > yMax {
				yMax = tail.Y
			}

			if tail.X < xMin {
				xMin = tail.X
			}
			if tail.Y < yMin {
				yMin = tail.Y
			}
		}
	}

	xOffset := Abs(xMin) + 1
	yOffset := Abs(yMin) + 1

	bridge := make([][]bool, xOffset+xMax, yOffset+yMax)

	for i := range bridge {
		bridge[i] = make([]bool, xOffset+yMax)
		for j := range bridge[i] {
			bridge[i][j] = false
		}
	}

	for i := range positions {
		curr := positions[i]
		bridge[curr.X+xOffset-1][curr.Y+yOffset-1] = true
	}

	sum := 0

	for i := range bridge {
		for j := range bridge[i] {
			if bridge[i][j] {
				sum++
			}
		}
	}

	fmt.Printf("Part 1: %v\n", sum)
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	head := Position{
		X: 0,
		Y: 0,
	}
	tail := Position{
		X: 0,
		Y: 0,
	}

	positions := []Position{}

	positions = append(positions, Position{X: 0, Y: 0})

	xMax := 0
	yMax := 0

	xMin := 0
	yMin := 0

	for scanner.Scan() {
		move := strings.Split(scanner.Text(), " ")
		direction := move[0]
		spaces, _ := strconv.Atoi(move[1])
		for i := 0; i < spaces; i++ {
			if direction == "U" {
				head.Y++
				if Abs(tail.Y-head.Y) > 1 && Abs(tail.X-head.X) > 0 {
					tail.X = head.X
				}
				if Abs(tail.Y-head.Y) > 1 {
					tail.Y++
				}
			} else if direction == "D" {
				head.Y--
				if Abs(tail.Y-head.Y) > 1 && Abs(tail.X-head.X) > 0 {
					tail.X = head.X
				}
				if Abs(tail.Y-head.Y) > 1 {
					tail.Y--
				}
			} else if direction == "L" {
				head.X--
				if Abs(tail.X-head.X) > 1 && Abs(tail.Y-head.Y) > 0 {
					tail.Y = head.Y
				}
				if Abs(tail.X-head.X) > 1 {
					tail.X--
				}
			} else if direction == "R" {
				head.X++
				if Abs(tail.X-head.X) > 1 && Abs(tail.Y-head.Y) > 0 {
					tail.Y = head.Y
				}
				if Abs(tail.X-head.X) > 1 {
					tail.X++
				}
			}
			positions = append(positions, Position{X: tail.X, Y: tail.Y})
			if tail.X > xMax {
				xMax = tail.X
			}
			if tail.Y > yMax {
				yMax = tail.Y
			}

			if tail.X < xMin {
				xMin = tail.X
			}
			if tail.Y < yMin {
				yMin = tail.Y
			}
		}
	}

	xOffset := Abs(xMin) + 1
	yOffset := Abs(yMin) + 1

	bridge := make([][]bool, xOffset+xMax, yOffset+yMax)

	for i := range bridge {
		bridge[i] = make([]bool, xOffset+yMax)
		for j := range bridge[i] {
			bridge[i][j] = false
		}
	}

	for i := range positions {
		curr := positions[i]
		bridge[curr.X+xOffset-1][curr.Y+yOffset-1] = true
	}

	sum := 0

	for i := range bridge {
		for j := range bridge[i] {
			if bridge[i][j] {
				sum++
			}
		}
	}

	fmt.Printf("Part 2: %v\n", sum)
}
