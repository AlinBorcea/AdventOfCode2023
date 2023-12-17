package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename = "puzzle.txt"

type NumCoord struct {
	num, x, y int
}

func main() {
	lines := readFile(filename)
	sum := partNumberSum(lines)

	fmt.Println(sum)
}

func readFile(filename string) (lines []string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func partNumberSum(lines []string) (sum int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] >= '0' && lines[i][j] <= '9' {
				numCoord := getNumCoord(lines[i], j)
				if adjacentToSymbol(lines, i, numCoord) {
					sum += numCoord.num
					j = numCoord.y
				}
			}
		}
	}
	return
}

func getNumCoord(str string, start int) NumCoord {
	n := 0
	i := start
	for ; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
		n = n*10 + int(str[i]-'0')
	}

	return NumCoord{num: n, x: start, y: i - 1}
}

func adjacentToSymbol(lines []string, i int, numCoord NumCoord) bool {
	//horizontal
	if numCoord.x > 0 && charIsSymbol(lines[i][numCoord.x-1]) {
		return true

	}
	if numCoord.y < len(lines[i])-1 && charIsSymbol(lines[i][numCoord.y+1]) {
		return true

	}
	if i > 0 { //vertical
		j := numCoord.x - 1
		z := numCoord.y + 1
		if numCoord.x == 0 {
			j = 0
		}

		if numCoord.y == len(lines[i-1])-1 {
			z = numCoord.y
		}

		for ; j <= z; j++ {
			if charIsSymbol(lines[i-1][j]) {
				return true
			}
		}

	}
	if i < len(lines)-1 {
		j := numCoord.x - 1
		z := numCoord.y + 1
		if numCoord.x == 0 {
			j = 0
		}

		if numCoord.y == len(lines[i+1])-1 {
			z = numCoord.y
		}

		for ; j <= z; j++ {
			if charIsSymbol(lines[i+1][j]) {
				return true
			}
		}
	}

	return false
}

func charIsSymbol(ch byte) bool {
	return ch != '.' && !(ch >= '0' && ch <= '9')
}
