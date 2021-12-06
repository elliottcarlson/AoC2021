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

type Node struct {
	Value int
	Count int
	Left  *Node
	Right *Node
}

func (n *Node) addEntry(entry string) {
	var parent *Node = n
	for _, bit := range entry {
		value, _ := strconv.Atoi(string(bit))
		if value == 0 {
			if parent.Left == nil {
				parent.Left = &Node{
					Value: 0,
					Count: 0,
				}
			}
			parent.Left.Count++
			parent = parent.Left
		} else {
			if parent.Right == nil {
				parent.Right = &Node{
					Value: 1,
					Count: 0,
				}
			}
			parent.Right.Count++
			parent = parent.Right
		}
	}
}

func (n *Node) getOxygen() int64 {
	var binary [dataLength]int
	var node *Node = n
	for i := 0; i < dataLength; i++ {
		if (node.Left != nil && node.Right == nil) || node.Left.Count > node.Right.Count {
			binary[i] = node.Left.Value
			node = node.Left
		} else {
			binary[i] = node.Right.Value
			node = node.Right
		}
	}

	return intArrayToDecimal(binary)
}

func (n *Node) getCO2() int64 {
	var binary [dataLength]int
	var node *Node = n
	for i := 0; i < dataLength; i++ {
		if node.Left != nil && node.Right != nil {
			if node.Right.Count < node.Left.Count {
				binary[i] = node.Right.Value
				node = node.Right
			} else {
				binary[i] = node.Left.Value
				node = node.Left
			}
		} else if node.Left == nil {
			binary[i] = node.Right.Value
			node = node.Right
		} else {
			binary[i] = node.Left.Value
			node = node.Left
		}
	}

	return intArrayToDecimal(binary)
}

func main() {
	root := Node{}

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		root.addEntry(current)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	oxygen := root.getOxygen()
	co2 := root.getCO2()

	fmt.Printf("Oxygen Generator Rating (%d) * CO2 scrubber rating (%d) = Life Support Rating (%d)\n", oxygen, co2, oxygen*co2)
}

func intArrayToDecimal(a [dataLength]int) int64 {
	binary := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ""), "[]")
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return output
}
