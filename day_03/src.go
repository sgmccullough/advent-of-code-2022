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

	sum := 0

	for scanner.Scan() {
		curr := scanner.Text()
		firstHalf := curr[0 : len(curr)/2]
		secondHalf := curr[len(curr)/2:]
		commonChar := findCommonChar(firstHalf, secondHalf)
		if int(commonChar) >= 96 {
			sum += int(commonChar) - 96
		} else if int(commonChar) <= 91 {
			sum += int(commonChar) - 38
		}
	}

	fmt.Printf("Part 1: %v\n", sum)
}

// Assumes that there is at least a common char
func findCommonChar(a string, b string) rune {
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			return c
		}
	}
	return ' '
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0

	for i := 0; i < len(lines); i += 3 {
		commonChar := findCommonChar_3(lines[i], lines[i+1], lines[i+2])
		if int(commonChar) >= 96 {
			sum += int(commonChar) - 96
		} else if int(commonChar) <= 91 {
			sum += int(commonChar) - 38
		}
	}

	fmt.Printf("Part 2: %v\n", sum)
}

// Assumes that there is at least a common char
func findCommonChar_3(a string, b string, c string) rune {
	for _, d := range a {
		if strings.Contains(b, string(d)) && strings.Contains(c, string(d)) {
			return d
		}
	}
	return ' '
}
