package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

const inputFile = "input"

type Boards struct {
    Board []Board
}

type Board struct {
    Squares  []Square
}

type Square struct {
    Value  int
    Called bool
}

func (b *Boards) Call(val int) (bool, int) {
    for i := range b.Board {
        b.Board[i].Call(val)
        if b.Board[i].isWinner() {
            unmarked := 0
            for j := range b.Board[i].Squares {
                if b.Board[i].Squares[j].Called == false {
                    unmarked = unmarked + b.Board[i].Squares[j].Value
                }
            }
            return true, unmarked * val
        }
    }

    return false, 0
}


func (b *Board) isWinner() bool {
    winningIndexes := [][]int{
        []int{0, 1, 2, 3, 4},
        []int{5, 6, 7, 8, 9},
        []int{10, 11, 12, 13, 14},
        []int{15, 16, 17, 18, 19},
        []int{20, 21, 22, 23, 24},
        []int{0, 5, 10, 15, 20},
        []int{1, 6, 11, 16, 21},
        []int{2, 7, 12, 17, 22},
        []int{3, 8, 13, 18, 23},
        []int{4, 9, 14, 19, 24},
//        []int{1, 7, 13, 19, 25},
//        []int{5, 9, 13, 17, 21},
    }

    for i := range winningIndexes {
        counter := 0
        for _, j := range winningIndexes[i] {
            if b.Squares[j].Called {
                counter++
            }
        }
        if counter == 5 {
            return true
        }
    }

    return false
}

func (b *Board) Call(val int) {
    for i := range b.Squares {
        if b.Squares[i].Value == val {
            b.Squares[i].Called = true
            return
        }
    }
}


func main() {
    file, err := os.Open(inputFile)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    numbers := strings.Split(scanner.Text(), ",")

    space := regexp.MustCompile(`\s+`)

    boards := Boards{}
    squares := []Square{}
    for scanner.Scan() {
        current := scanner.Text()
        if current == "" {
            if len(squares) == 25 {
                boards.Board = append(boards.Board, Board{
                    Squares: squares,
                })
            }
            squares = []Square{}
            continue
        }
        trimmed := strings.TrimSpace(space.ReplaceAllString(current, " "))
        tiles := strings.Split(trimmed, " ")
        for _, tile := range tiles {
            val, _ := strconv.Atoi(tile)
            squares = append(squares, Square{
                Value: val,
                Called: false,
            })
        }
    }
    boards.Board = append(boards.Board, Board{
        Squares: squares,
    })

    for _, num := range numbers {
        val, _ := strconv.Atoi(num)
        winner, b := boards.Call(val)
        if winner {
            fmt.Printf("Winner found! Board %d\n", b)
            break
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
