package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const filename = "puzzle.txt"

func main() {
	/*
		[]string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		}
	*/
	lines := readPuzzle(filename)
	lineValues := findLineValues(lines)
	sum := sumOfLineValues(lineValues)

	fmt.Println(sum)
}

func readPuzzle(filename string) []string {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func findLineValues(lines []string) []int {
	values := make([]int, len(lines))

	for i, line := range lines {
		x, y := findLinePairs(line)
		if x != -1 && y != -1 {
			values[i] = x*10 + y
		} else {
			values[i] = -1
		}
	}

	return values
}

func findLinePairs(line string) (int, int) {
	x, y := -1, -1

	for _, c := range line {
		if unicode.IsDigit(c) {
			if x == -1 {
				x = int(c - '0')
				y = x
			} else {
				y = int(c - '0')
			}
		}
	}

	return x, y
}

func sumOfLineValues(vals []int) int {
	sum := 0

	for _, x := range vals {
		if x < 0 {
			return -1
		}
		sum += x
	}

	return sum
}
