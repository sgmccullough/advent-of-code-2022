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
	part1_borked()

	part1()

	part2()
}

type Tree struct {
	Height int
	Above  *Tree
	Below  *Tree
	Left   *Tree
	Right  *Tree
}

func part1() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var forest [][]int

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, curr := range line {
			row[i] = int(curr)
		}
		forest = append(forest, row)
	}

	highest := [4][]int{make([]int, len(forest[0])), make([]int, len(forest)), make([]int, len(forest[0])), make([]int, len(forest))}

	// init
	for i := range highest {
		for j := range highest[i] {
			highest[i][j] = -1
		}
	}

	isVisible := make([][]bool, len(forest))
	for i := range isVisible {
		isVisible[i] = make([]bool, len(forest[0]))
	}

	sum := 0

	walkSummary := func(row, column int, highest *int) {
		if forest[row][column] > *highest {
			if *highest = forest[row][column]; !isVisible[row][column] {
				isVisible[row][column] = true
				sum++
			}
		}
	}

	for down := range forest {
		up := len(forest) - down - 1
		for right := range forest[down] {
			left := len(forest[down]) - right - 1
			walkSummary(down, right, &highest[0][right])
			walkSummary(down, right, &highest[1][down])
			walkSummary(up, left, &highest[2][left])
			walkSummary(up, left, &highest[3][up])
		}
	}

	fmt.Printf("Part 1: %v\n", sum)
}

func part1_borked() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var forest [][]Tree

	for scanner.Scan() {
		var row []Tree
		for _, elem := range strings.Split(scanner.Text(), "") {
			n, _ := strconv.Atoi(elem)
			tree := Tree{
				Height: n,
				Above:  nil,
				Below:  nil,
				Left:   nil,
				Right:  nil,
			}
			row = append(row, tree)
		}
		forest = append(forest, row)
	}

	// i is for column, j is for row
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			// check above
			if i != 0 {
				forest[i][j].Above = &forest[i-1][j]
			}
			// check below
			if i != len(forest)-1 {
				forest[i][j].Below = &forest[i+1][j]
			}
			// check left
			if j != 0 {
				forest[i][j].Left = &forest[i][j-1]
			}
			// check right
			if j != len(forest[i])-1 {
				forest[i][j].Right = &forest[i][j+1]
			}
		}
	}

	sum := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			if isVisible(forest[i][j], forest[i][j].Height, "above") ||
				isVisible(forest[i][j], forest[i][j].Height, "below") ||
				isVisible(forest[i][j], forest[i][j].Height, "left") ||
				isVisible(forest[i][j], forest[i][j].Height, "right") {
				sum++
			}
		}
	}

	fmt.Printf("Part 1 (Borked): %v\n", sum)
}

func isVisible(tree Tree, max int, direction string) bool {
	if direction == "above" {
		if tree.Above == nil {
			return true
		} else if tree.Above.Height <= tree.Height && tree.Above.Height < max {
			return isVisible(*tree.Above, max, direction)
		}
	} else if direction == "below" {
		if tree.Below == nil {
			return true
		} else if tree.Below.Height <= tree.Height && tree.Below.Height < max {
			return isVisible(*tree.Below, max, direction)
		}
	} else if direction == "left" {
		if tree.Left == nil {
			return true
		} else if tree.Left.Height <= tree.Height && tree.Left.Height < max {
			return isVisible(*tree.Left, max, direction)
		}
	} else if direction == "right" {
		if tree.Right == nil {
			return true
		} else if tree.Right.Height <= tree.Height && tree.Right.Height < max {
			return isVisible(*tree.Right, max, direction)
		}
	}
	return false
}

func part2() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var forest [][]int

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, curr := range line {
			row[i] = int(curr)
		}
		forest = append(forest, row)
	}

	var highest int

	for row := 1; row < len(forest)-1; row++ {
		for column := 1; column < len(forest[row])-1; column++ {
			up := row - 1
			down := row + 1
			left := column - 1
			right := column + 1

			for up > 0 && forest[up][column] < forest[row][column] {
				up--
			}
			for down < len(forest)-1 && forest[down][column] < forest[row][column] {
				down++
			}
			for left > 0 && forest[row][left] < forest[row][column] {
				left--
			}
			for right < len(forest[row])-1 && forest[row][right] < forest[row][column] {
				right++
			}

			sum := (row - up) * (down - row) * (column - left) * (right - column)

			if sum > highest {
				highest = sum
			}
		}
	}

	fmt.Printf("Part 2: %v\n", highest)
}
