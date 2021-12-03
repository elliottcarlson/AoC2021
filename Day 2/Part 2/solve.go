package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var depth int
    var horizontal int
    var aim int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")
        val, _ := strconv.Atoi(s[1])
        switch (s[0]) {
            case "forward":
                horizontal = horizontal + val
                depth = depth + (aim * val)
            case "up":
                aim = aim - val
            case "down":
                aim = aim + val
        }
    }

    fmt.Printf("Horizontal: %d\n", horizontal)
    fmt.Printf("Depth: %d\n", depth)
    fmt.Printf("Result: %d\n", (depth * horizontal))

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
