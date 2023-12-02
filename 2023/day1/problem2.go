package main

import (
    "bufio"
	"fmt"
    "os"
    "unicode"
    "strconv"
)

func read_stdin() []string {
    scanner := bufio.NewScanner(os.Stdin)
    var input []string
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
    return input
}

func main() {
    sum := 0
    tmpVal := 0
    strNumbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    for _, line := range read_stdin() {
        lineVals := []string{}
        for i, char := range line {
            if unicode.IsNumber(char) {
                lineVals = append(lineVals, string(char))
                continue
            }
            for j, strNumber := range strNumbers {
                if len(line[i:]) < len(strNumber) {
                    continue
                }
                if strNumber == line[i:i+len(strNumber)] {
                    lineVals = append(lineVals, strconv.Itoa(j))
                    break
                }
            }
        }
        tmpVal, _ = strconv.Atoi(lineVals[0] + lineVals[len(lineVals)-1])
        sum = sum + tmpVal
    }
    fmt.Println(sum)
}
