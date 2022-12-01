package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	largest := 0
	current := 0

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if current > largest {
				largest = current
			}
			current = 0
		} else {
			current += curr
		}
	}

	fmt.Printf("Part 1: %v\n", largest)
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	current := 0
	var vals []int

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			vals = append(vals, current)
			current = 0
		} else {
			current += curr
		}
	}

	sort.Ints(vals)

	length := len(vals)

	output := vals[length-1] + vals[length-2] + vals[length-3]

	fmt.Printf("Part 2: %v\n", output)
}
