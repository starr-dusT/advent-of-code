package main

import (
	L "day3/lib"
	"fmt"
)

type blockNumber struct {
    number int
    block [][2]int
}

func main() {
    sum := 0
    gearDict := make(map[[2]int][]int)

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
        lengthLine := len(input[blockIndexes[i][0][1]])
        adjacentIndexes := L.Adjacent_indexes(indexes, lengthLine, len(input))
        for _, index := range adjacentIndexes {
            char := rune(input[index[1]][index[0]])
            if string(char) == "*" {
                gearDict[index] = append(gearDict[index], blockNumbers[i])
                break
            }
        }
    }

    for _, numbers := range gearDict {
        tmpNumber := 1
        if len(numbers) > 1 {
            for _, number := range numbers {
                tmpNumber = tmpNumber * number
            }
            sum = sum + tmpNumber
        }
    }
    fmt.Println(sum)
}
