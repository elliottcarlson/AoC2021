package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input"
const dayCycle = 256

type Fish struct {
	Age [9]int
}

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := scanner.Text()
	seeds := strings.Split(data, ",")

	school := Fish{}

	for _, seed := range seeds {
		val, _ := strconv.Atoi(seed)
		school.Age[val]++
	}

	for day := 1; day <= dayCycle; day++ {
		mutation := Fish{}
		for age, count := range school.Age {
			if age == 0 {
				mutation.Age[6] += count
				mutation.Age[8] += count
			} else {
				mutation.Age[age-1] += count
			}
		}
		school = mutation
	}

	total := 0
	for _, count := range school.Age {
		total += count
	}

	fmt.Printf("Total fish: %d\n", total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
