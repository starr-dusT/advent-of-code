package main

import (
	L "github.com/starr-dusT/advent-of-code/lib"
	"fmt"
	"unicode"
)

type blockNumber struct {
    number int
    block [][2]int
}

func main() {
    sum := 0
    
    //get numbers and indexes where numbers are in grid
    blockNumbers := []int{}
    blockIndexes := [][][2]int{}
    input := []string{}
    for i, line := range L.Read_stdin() {
        input = append(input, line)
        numbers, indexes := L.FindNumbers(line)
        for j, _ := range numbers {
            blockNumbers = append(blockNumbers, numbers[j])
            blockIndexes = append(blockIndexes, L.MakeBlockFromLine(i, indexes[j]))
        }
    }

    //determine adjacent characeters
    for i, indexes := range blockIndexes {
        fmt.Println(blockNumbers[i])
        lengthLine := len(input[blockIndexes[i][0][1]])
        adjacentIndexes := L.Adjacent_indexes(indexes, lengthLine, len(input))
        fmt.Println(adjacentIndexes)
        for _, index := range adjacentIndexes {
            char := rune(input[index[1]][index[0]])
            fmt.Print(string(char))
            if !unicode.IsNumber(char) && string(char) != "." {
                sum = sum + blockNumbers[i]
                fmt.Print("Found")
                break
            }
        }
        fmt.Println("")
    }
    fmt.Println(sum)
}
