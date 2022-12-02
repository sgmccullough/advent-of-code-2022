package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	score := 0

	for scanner.Scan() {
		curr := strings.Fields(scanner.Text())
		if curr[1] == "X" { // Rock
			score += 1
		} else if curr[1] == "Y" { // Paper
			score += 2
		} else if curr[1] == "Z" { // Scissors
			score += 3
		}

		if (curr[0] == "A" && curr[1] == "Y") || (curr[0] == "B" && curr[1] == "Z") || (curr[0] == "C" && curr[1] == "X") { // Win
			score += 6
		} else if (curr[0] == "A" && curr[1] == "X") || (curr[0] == "B" && curr[1] == "Y") || (curr[0] == "C" && curr[1] == "Z") { // Draw
			score += 3
		}
	}

	fmt.Printf("Part 1: %v\n", score)
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	score := 0

	for scanner.Scan() {
		curr := strings.Fields(scanner.Text())
		if curr[1] == "X" { // Lose
			if curr[0] == "A" {
				score += 3
			} else if curr[0] == "B" {
				score += 1
			} else if curr[0] == "C" {
				score += 2
			}
		} else if curr[1] == "Y" { // Draw
			score += 3
			if curr[0] == "A" {
				score += 1
			} else if curr[0] == "B" {
				score += 2
			} else if curr[0] == "C" {
				score += 3
			}
		} else if curr[1] == "Z" { // Win
			score += 6
			if curr[0] == "A" {
				score += 2
			} else if curr[0] == "B" {
				score += 3
			} else if curr[0] == "C" {
				score += 1
			}
		}
	}

	fmt.Printf("Part 2: %v\n", score)
}
