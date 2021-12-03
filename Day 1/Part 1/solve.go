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

    var last int
    increases := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        current, _ := strconv.Atoi(scanner.Text())
        if last > 0 {
            if current > last {
                increases = increases + 1
            }
        }
        last, _ = strconv.Atoi(scanner.Text())
    }

    fmt.Println(increases)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
