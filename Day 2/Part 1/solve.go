package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const inputFile = "input"

func main() {
    file, err := os.Open(inputFile)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var depth int
    var horizontal int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")
        val, _ := strconv.Atoi(s[1])
        switch (s[0]) {
            case "forward":
                horizontal = horizontal + val
            case "up":
                depth = depth - val
            case "down":
                depth = depth + val
        }
    }

    fmt.Printf("Horizontal Position: %d\n", horizontal)
    fmt.Printf("Depth Position: %d\n", depth)
    fmt.Printf("Answer: %d\n", depth * horizontal)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
