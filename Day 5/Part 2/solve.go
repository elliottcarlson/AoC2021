package main

import (
	"bufio"
	"fmt"
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

		dx := 0
		dy := 0

		if x2 > x1 {
			dx = 1
		} else if x1 > x2 {
			dx = -1
		}

		if y2 > y1 {
			dy = 1
		} else if y1 > y2 {
			dy = -1
		}

		x := x1
		y := y1

		for {
			grid.Columns[y].Rows[x].Count++
			if x == x2 && y == y2 {
				break
			}
			x += dx
			y += dy
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

	fmt.Printf("Intersections: %d\n", count)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
