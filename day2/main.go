package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//const infile = "in.txt"

const infile = "puzzle.txt"

func bagMap() map[string]int {
	return map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

type Record struct {
	id     int
	blocks []map[string]int
}

func main() {
	inLines, err := readLines(infile)
	if err != nil {
		panic(err.Error())
	}

	records := createRecords(inLines)
	sum := idSumOfPossibleGames(records)

	//fmt.Println(records)
	fmt.Println(sum)
}

func readLines(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(infile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func createRecords(lines []string) []Record {
	var records []Record

	for _, line := range lines {
		records = append(records, createRecord(line))
	}

	return records
}

func createRecord(line string) (r Record) {
	gameIdSet := strings.Split(line, ": ")
	gameId := strings.Split(gameIdSet[0], " ")
	x, _ := strconv.ParseInt(gameId[1], 0, 32)
	r.id = int(x)

	set := strings.Split(gameIdSet[1], "; ")

	for _, val := range set {
		blocks := strings.Split(val, ", ")
		m := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, block := range blocks {
			valKey := strings.Split(block, " ")
			val1, _ := strconv.ParseInt(valKey[0], 0, 32)
			val := int(val1)
			key := valKey[1]
			m[key] += val
		}

		r.blocks = append(r.blocks, m)
	}

	return
}

func idSumOfPossibleGames(records []Record) int {
	sum := 0
	blockMap := bagMap()

	for _, rec := range records {
		if validRecord(rec, blockMap) {
			sum += rec.id
		}
	}

	return sum
}

func validRecord(rec Record, blockMap map[string]int) bool {
	for _, block := range rec.blocks {

		for key := range block {
			if block[key] > blockMap[key] {
				return false
			}
		}
	}
	return true
}
