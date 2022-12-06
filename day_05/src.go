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

	pos := 0
	stackNumber := 0
	var vals [2][]string

	for scanner.Scan() {
		curr := scanner.Text()
		if curr == "" {
			pos = 1
		} else if curr[1] == '1' {
			stackNumber = (len(curr) + 1) / 4
		} else {
			vals[pos] = append(vals[pos], curr)
		}
	}

	cargo := make([][]string, stackNumber)

	for _, s := range vals[0] {
		for j := 1; j < len(s); j += 4 {
			if s[j] != ' ' {
				cargo[(j-1)/4] = append([]string{string(s[j])}, cargo[(j-1)/4]...)
			}
		}
	}

	for _, s := range vals[1] {
		curr := strings.Fields(s)
		toMove, _ := strconv.Atoi(curr[1])
		from, _ := strconv.Atoi(curr[3])
		to, _ := strconv.Atoi(curr[5])
		for i := 1; i <= toMove; i++ {
			temp := cargo[from-1][len(cargo[from-1])-1]
			cargo[from-1] = cargo[from-1][:len(cargo[from-1])-1]
			cargo[to-1] = append(cargo[to-1], string(temp))
		}
	}

	output := ""

	for i := 0; i < len(cargo); i++ {
		output = output + cargo[i][len(cargo[i])-1]
	}

	fmt.Printf("Part 1: %s\n", output)
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	pos := 0
	stackNumber := 0
	var vals [2][]string

	for scanner.Scan() {
		curr := scanner.Text()
		if curr == "" {
			pos = 1
		} else if curr[1] == '1' {
			stackNumber = (len(curr) + 1) / 4
		} else {
			vals[pos] = append(vals[pos], curr)
		}
	}

	cargo := make([][]string, stackNumber)

	for _, s := range vals[0] {
		for j := 1; j < len(s); j += 4 {
			if s[j] != ' ' {
				cargo[(j-1)/4] = append([]string{string(s[j])}, cargo[(j-1)/4]...)
			}
		}
	}

	for _, s := range vals[1] {
		curr := strings.Fields(s)
		toMove, _ := strconv.Atoi(curr[1])
		from, _ := strconv.Atoi(curr[3])
		to, _ := strconv.Atoi(curr[5])
		temp := cargo[from-1][len(cargo[from-1])-(toMove) : len(cargo[from-1])]
		cargo[from-1] = cargo[from-1][:len(cargo[from-1])-toMove]
		cargo[to-1] = append(cargo[to-1], temp...)
	}

	output := ""

	for i := 0; i < len(cargo); i++ {
		output = output + cargo[i][len(cargo[i])-1]
	}

	fmt.Printf("Part 2: %s\n", output)
}
