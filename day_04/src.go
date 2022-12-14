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

	sum := 0

	for scanner.Scan() {
		curr := scanner.Text()
		removeComma := strings.Split(curr, ",")
		first := strings.Split(removeComma[0], "-")
		second := strings.Split(removeComma[1], "-")
		one_one, _ := strconv.Atoi(first[0])
		one_two, _ := strconv.Atoi(first[1])
		two_one, _ := strconv.Atoi(second[0])
		two_two, _ := strconv.Atoi(second[1])

		if one_one <= two_one && one_two >= two_two {
			sum += 1
		} else if two_one <= one_one && two_two >= one_two {
			sum += 1
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

	sum := 0

	for scanner.Scan() {
		curr := scanner.Text()
		removeComma := strings.Split(curr, ",")
		first := strings.Split(removeComma[0], "-")
		second := strings.Split(removeComma[1], "-")
		one_one, _ := strconv.Atoi(first[0])
		one_two, _ := strconv.Atoi(first[1])
		two_one, _ := strconv.Atoi(second[0])
		two_two, _ := strconv.Atoi(second[1])

		// a-b, c-d
		// a <= c && b >= c
		// c <= a && c >= b
		if (one_one <= two_one && one_two >= two_one) || (two_one <= one_one && two_one >= one_two) {
			sum += 1
		} else if one_one <= two_two && one_two >= two_two {
			sum += 1
		} else if two_one <= one_one && two_two >= one_one {
			sum += 1
		}
	}

	fmt.Printf("Part 2: %v\n", sum)
}
