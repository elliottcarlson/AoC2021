package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const gridSize = 999      // 10
const inputFile = "input" // input.test

type Grid struct {
	Columns [gridSize]Column
}

type Column struct {
	Rows [gridSize]Row
}

type Row struct {
	Count int
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := Grid{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()

		coords := strings.Split(current, " -> ")
		first := strings.Split(coords[0], ",")
		second := strings.Split(coords[1], ",")

		x1, _ := strconv.Atoi(first[0])
		y1, _ := strconv.Atoi(first[1])
		x2, _ := strconv.Atoi(second[0])
		y2, _ := strconv.Atoi(second[1])

		if x1 == x2 || y1 == y2 {
			for _, x := range Range(x1, x2) {
				for _, y := range Range(y1, y2) {
					grid.Columns[y].Rows[x].Count++
				}
			}
		}
	}

	count := 0
	for _, column := range grid.Columns {
		for _, row := range column.Rows {
			if row.Count > 1 {
				count++
			}
		}
	}

	fmt.Printf("%d\n", count)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func Range(start, end int) []int {
	if start == end {
		return []int{start}
	}

	count := int(math.Abs(float64(start)-float64(end))) + 1

	step := 1
	if start > end {
		step = -1
	}

	s := make([]int, count)

	for i := range s {
		s[i] = start
		start += step
	}

	return s
}
