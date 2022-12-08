package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// This solution is really not great, but I don't have the willpower to figure out how to do this more effectively lol
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

	scanner.Scan()
	input := scanner.Text()

	var buffer []byte
	buffer = append(buffer, input[0], input[1], input[2], input[3])

	pos := 4
	for !(is_unique(buffer[len(buffer)-1], buffer[:len(buffer)-1]) && is_unique(buffer[len(buffer)-2], buffer[:len(buffer)-2]) && is_unique(buffer[len(buffer)-3], buffer[:len(buffer)-3])) {
		buffer = append(buffer, input[pos])
		buffer = buffer[1:5]

		pos++
	}

	fmt.Printf("Part 1: %v\n", pos)
}

func is_unique(val byte, input []byte) bool {
	for _, elem := range input {
		if elem == val {
			return false
		}
	}
	return true
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	input := scanner.Text()

	var buffer []byte
	buffer = append(buffer, input[0], input[1], input[2], input[3], input[4], input[5], input[6], input[7], input[8], input[9], input[10], input[11], input[12], input[13])

	pos := 14
	for !(is_unique(buffer[len(buffer)-1], buffer[:len(buffer)-1]) && is_unique(buffer[len(buffer)-2], buffer[:len(buffer)-2]) && is_unique(buffer[len(buffer)-3], buffer[:len(buffer)-3]) && is_unique(buffer[len(buffer)-4], buffer[:len(buffer)-4]) && is_unique(buffer[len(buffer)-5], buffer[:len(buffer)-5]) && is_unique(buffer[len(buffer)-6], buffer[:len(buffer)-6]) && is_unique(buffer[len(buffer)-7], buffer[:len(buffer)-7]) && is_unique(buffer[len(buffer)-8], buffer[:len(buffer)-8]) && is_unique(buffer[len(buffer)-9], buffer[:len(buffer)-9]) && is_unique(buffer[len(buffer)-10], buffer[:len(buffer)-10]) && is_unique(buffer[len(buffer)-11], buffer[:len(buffer)-11]) && is_unique(buffer[len(buffer)-12], buffer[:len(buffer)-12]) && is_unique(buffer[len(buffer)-13], buffer[:len(buffer)-13])) {
		buffer = append(buffer, input[pos])
		buffer = buffer[1:15]

		pos++
	}

	fmt.Printf("Part 2: %v\n", pos)
}
