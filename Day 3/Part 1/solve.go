package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dataLength = 12
const inputFile = "input"

type Values struct {
	Zero    [dataLength]int
	One     [dataLength]int
	Gamma   [dataLength]int
	Epsilon [dataLength]int
}

func main() {
	output := Values{}

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		for i, value := range current {
			if fmt.Sprintf("%c", value) == "0" {
				output.Zero[i] = output.Zero[i] + 1
			} else {
				output.One[i] = output.One[i] + 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := range output.Zero {
		if output.Zero[i] > output.One[i] {
			output.Gamma[i] = 0
			output.Epsilon[i] = 1
		} else {
			output.Gamma[i] = 1
			output.Epsilon[i] = 0
		}
	}

	gamma := intArrayToDecimal(output.Gamma)
	epsilon := intArrayToDecimal(output.Epsilon)

	fmt.Printf("Gamma Rate (%d) * Epsilon Rate (%d) = Power Consumption (%d)\n", gamma, epsilon, gamma*epsilon)
}

func intArrayToDecimal(a [dataLength]int) int64 {
	binary := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ""), "[]")
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return output
}
