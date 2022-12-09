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

type Directory struct {
	Name            string
	Size            int
	ParentDirectory *Directory
	SubDirectories  []*Directory
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
	initialize := strings.Fields(scanner.Text())

	directory := Directory{
		Name:            initialize[2],
		Size:            0,
		ParentDirectory: nil,
		SubDirectories:  []*Directory{},
	}

	build(scanner, &directory)

	sum := walk_sum(&directory)

	fmt.Printf("Part 1: %v\n", sum)
}

func build(scanner *bufio.Scanner, current *Directory) {
	currentSize := 0
	level := 0

	for scanner.Scan() {
		curr := strings.Fields(scanner.Text())
		if curr[0] == "$" && curr[1] == "cd" { // Traversing a Directory
			if curr[2] == ".." {
				currentSize = current.Size
				current = current.ParentDirectory
				current.Size = current.Size + currentSize // Update the current size to include the last sub dir
				level--
			} else {
				index := sort.Search(len(current.SubDirectories), func(i int) bool {
					return string(current.SubDirectories[i].Name) >= curr[2]
				})
				current = current.SubDirectories[index]
				level++
			}
		} else if curr[0] == "dir" {
			newDirectory := Directory{
				Name:            curr[1],
				Size:            0,
				ParentDirectory: current,
				SubDirectories:  []*Directory{},
			}
			current.SubDirectories = append(current.SubDirectories, &newDirectory)
		} else {
			size, _ := strconv.Atoi(curr[0])
			current.Size = current.Size + size
		}
	}

	for i := 0; i < level; i++ { // reset us to the root of the tree
		currentSize = current.Size
		current = current.ParentDirectory
		current.Size = current.Size + currentSize
	}
}

func walk_sum(directory *Directory) int {
	size := 0
	if len(directory.SubDirectories) == 0 {
		if directory.Size <= 100000 {
			size += directory.Size
		}
	} else {
		for _, element := range directory.SubDirectories {
			if directory.Size <= 100000 {
				size += directory.Size
			}
			size += walk_sum(element)
		}
	}
	return size
}

func walk_min(directory *Directory, upperBound int, lowerBound int) int {
	if directory.Size >= lowerBound && directory.Size < upperBound {
		upperBound = directory.Size
	}
	for _, element := range directory.SubDirectories {
		bound := walk_min(element, upperBound, lowerBound)
		if bound >= lowerBound && bound < upperBound {
			upperBound = bound
		}
	}
	return upperBound
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
	initialize := strings.Fields(scanner.Text())

	directory := Directory{
		Name:            initialize[2],
		Size:            0,
		ParentDirectory: nil,
		SubDirectories:  []*Directory{},
	}

	build(scanner, &directory)

	maxSpace := 70000000
	reqSpace := 30000000
	unusedSpace := maxSpace - directory.Size
	minDirSize := reqSpace - unusedSpace

	minFileSize := walk_min(&directory, directory.Size, minDirSize)

	fmt.Printf("Part 2: %v\n", minFileSize)
}
