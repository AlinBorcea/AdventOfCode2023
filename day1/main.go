package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	/*lines := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}*/
	lines := readPuzzle(filename)
	lines = trimPuzzle(lines)
	lineValues := findLineValues(lines)
	sum := sumOfLineValues(lineValues)

	//fmt.Println(lines)
	//fmt.Println(lineValues)
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

func trimPuzzle(lines []string) []string {
	for i, line := range lines {
		lines[i] = trimLine(line)
	}
	return lines
}

func trimLine(line string) string {

	line = strings.ReplaceAll(line, "zerone", "01")
	line = strings.ReplaceAll(line, "oneight", "18")
	line = strings.ReplaceAll(line, "twone", "21")
	line = strings.ReplaceAll(line, "threeight", "38")
	line = strings.ReplaceAll(line, "fiveight", "58")
	line = strings.ReplaceAll(line, "sevenine", "79")
	line = strings.ReplaceAll(line, "eightwo", "82")
	line = strings.ReplaceAll(line, "nineight", "98")

	line = strings.ReplaceAll(line, "one", "1")
	line = strings.ReplaceAll(line, "two", "2")
	line = strings.ReplaceAll(line, "six", "6")

	line = strings.ReplaceAll(line, "zero", "0")
	line = strings.ReplaceAll(line, "four", "4")
	line = strings.ReplaceAll(line, "five", "5")
	line = strings.ReplaceAll(line, "nine", "9")

	line = strings.ReplaceAll(line, "three", "3")
	line = strings.ReplaceAll(line, "seven", "7")
	line = strings.ReplaceAll(line, "eight", "8")

	return line
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
