package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFile = "input"

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var last0 int
	var last1 int
	var last2 int
	var last_sum int
	increases := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())
		if last0 > 0 {
			if (last1 + last2 + current) > last_sum {
				increases = increases + 1
			}
		}
		last0 = last1
		last1 = last2
		last2, _ = strconv.Atoi(scanner.Text())
		last_sum = last0 + last1 + last2
	}

	fmt.Println(increases)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
