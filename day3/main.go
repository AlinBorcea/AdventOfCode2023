package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filename = "puzzle.txt"

func main() {
	lines := readFile(filename)
	lineLength := findLineLength(filename)
	sum := partNumberSum(lines, lineLength)

	fmt.Println(sum)
}

func readFile(filename string) string {
	buf, _ := os.ReadFile(filename)
	content := string(buf)
	content = strings.ReplaceAll(content, "\r\n", "")
	return content
}

func findLineLength(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return len(scanner.Text())
}

func partNumberSum(lines string, lineLength int) int {
	var sum, n, x, y int

	for i := 0; i < len(lines); i++ {
		if lines[i] >= '0' && lines[i] <= '9' {
			n, x, y = findNumbersCoordinates(lines, lineLength, i)
			if partNumber(lines, lineLength, x, y) {
				sum += n
			}
			i = y
		}
	}

	return sum
}

func findNumbersCoordinates(lines string, lineLength, index int) (int, int, int) {
	n := 0
	i := index
	for ; i%lineLength < lineLength && lines[i] >= '0' && lines[i] <= '9'; i++ {
		n = n*10 + int(lines[i]-'0')
	}

	return n, index, i - 1
}

func partNumber(lines string, lineLength, x, y int) bool {
	var x1, y1 int

	if x%lineLength == 0 {
		x1 = x
	} else {
		x1 = x - 1
	}
	if y%lineLength == lineLength-1 {
		y1 = y
	} else {
		y1 = y + 1
	}

	if subStringHasSymbol(lines[x1 : y1+1]) {
		return true
	}

	if x1-lineLength >= 0 && subStringHasSymbol(lines[x1-lineLength:y1-lineLength+1]) {
		return true
	}

	if x1+lineLength < len(lines) && subStringHasSymbol(lines[x1+lineLength:y1+lineLength+1]) {
		return true
	}

	return false
}

func subStringHasSymbol(str string) bool {
	for _, ch := range str {
		if charIsSymbol(byte(ch)) {
			return true
		}
	}
	return false
}

func charIsSymbol(ch byte) bool {
	return ch != '.' && !(ch >= '0' && ch <= '9')
}
