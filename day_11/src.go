package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()

	part2()
}

type Monkey struct {
	ID          int
	Items       []int
	Operation   []string // 0: 1st term, 1: operation, 2: 2nd term
	Test        int
	IfTrue      int
	IfFalse     int
	Inspections int
}

func part1() {
	Monkeys := initialize()

	rounds := 20

	for i := 0; i < rounds; i++ {
		for _, monkey := range Monkeys {
			for k, item := range monkey.Items {
				// count the inspection
				Monkeys[monkey.ID].Inspections++
				// operation
				first := 0
				second := 0
				if monkey.Operation[0] == "old" {
					first = item
				} else {
					first, _ = strconv.Atoi(monkey.Operation[0])
				}
				if monkey.Operation[2] == "old" {
					second = item
				} else {
					second, _ = strconv.Atoi(monkey.Operation[2])
				}

				total := 0
				if monkey.Operation[1] == "+" {
					total = first + second
				} else if monkey.Operation[1] == "-" {
					total = first - second
				} else if monkey.Operation[1] == "*" {
					total = first * second
				} else if monkey.Operation[1] == "/" {
					total = int(float64(first) / float64(second))
				}

				// bored
				total = total / 3

				// test
				if total%monkey.Test == 0 { // true
					Monkeys[monkey.IfTrue].Items = append(Monkeys[monkey.IfTrue].Items, total)
				} else {
					Monkeys[monkey.IfFalse].Items = append(Monkeys[monkey.IfFalse].Items, total)
				}
				Monkeys[monkey.ID].Items = remove(Monkeys[monkey.ID].Items, k)
			}
		}
	}

	inspections := make([]int, len(Monkeys))
	fmt.Println("Part 1:")
	for _, s := range Monkeys {
		fmt.Printf("Monkey %v inspected items %v times.\n", s.ID, s.Inspections)
		inspections = append(inspections, s.Inspections)
	}
	sort.Ints(inspections)
	fmt.Printf("Monkey Business: %v\n", inspections[len(inspections)-1]*inspections[len(inspections)-2])
}

func part2() {
	Monkeys := initialize()

	rounds := 100

	for i := 0; i < rounds; i++ {
		for _, monkey := range Monkeys {
			for k, item := range monkey.Items {
				// count the inspection
				Monkeys[monkey.ID].Inspections++
				// operation
				first := 0
				second := 0
				if monkey.Operation[0] == "old" {
					first = item
				} else {
					first, _ = strconv.Atoi(monkey.Operation[0])
				}
				if monkey.Operation[2] == "old" {
					second = item
				} else {
					second, _ = strconv.Atoi(monkey.Operation[2])
				}

				total := 0
				if monkey.Operation[1] == "+" {
					total = first + second
				} else if monkey.Operation[1] == "-" {
					total = first - second
				} else if monkey.Operation[1] == "*" {
					total = first * second
				} else if monkey.Operation[1] == "/" {
					total = int(float64(first) / float64(second))
				}

				// test
				if total%monkey.Test == 0 { // true
					Monkeys[monkey.IfTrue].Items = append(Monkeys[monkey.IfTrue].Items, total)
				} else {
					Monkeys[monkey.IfFalse].Items = append(Monkeys[monkey.IfFalse].Items, total)
				}
				Monkeys[monkey.ID].Items = remove(Monkeys[monkey.ID].Items, k)
			}
		}
	}

	inspections := make([]int, len(Monkeys))
	fmt.Println("Part 2:")
	for _, s := range Monkeys {
		fmt.Printf("Monkey %v inspected items %v times.\n", s.ID, s.Inspections)
		inspections = append(inspections, s.Inspections)
	}
	sort.Ints(inspections)
	fmt.Printf("Monkey Business: %v\n", inspections[len(inspections)-1]*inspections[len(inspections)-2])
}

func remove(s []int, index int) []int {
	if index >= len(s) {
		return s[:len(s)-1]
	}
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func initialize() []Monkey {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	currMonkey := 0
	var Monkeys []Monkey

	// init
	for scanner.Scan() {
		line := scanner.Text()
		curr := strings.Split(strings.TrimLeft(line, " "), " ")
		if curr[0] == "Starting" {
			for _, s := range curr[2:] {
				v, _ := strconv.Atoi(strings.Trim(s, ","))
				Monkeys[currMonkey].Items = append(Monkeys[currMonkey].Items, v)
			}
		} else if curr[0] == "Operation:" {
			Monkeys[currMonkey].Operation = []string{curr[3], curr[4], curr[5]}
		} else if curr[0] == "Test:" {
			Monkeys[currMonkey].Test, _ = strconv.Atoi(curr[3])
		} else if curr[0] == "If" {
			if curr[1] == "true:" {
				Monkeys[currMonkey].IfTrue, _ = strconv.Atoi(curr[5])
			} else if curr[1] == "false:" {
				Monkeys[currMonkey].IfFalse, _ = strconv.Atoi(curr[5])
			}
		} else if curr[0] == "Monkey" {
			currMonkey = int(curr[1][0] - '0')
			newMonkey := Monkey{
				ID:          currMonkey,
				Inspections: 0,
			}
			Monkeys = append(Monkeys, newMonkey)
		}
	}
	return Monkeys
}
