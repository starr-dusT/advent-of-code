package main

import (
    "bufio"
	"fmt"
  	"strings"
    "os"
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
    abcs := "abcdefghijklmnopqrstuvwxyz"
    sum := 0
    tmpVal := 0
    for _, line := range read_stdin() {
        trimLine := strings.Trim(line, abcs)
        tmpVal, _ = strconv.Atoi(string(trimLine[0]) + string(trimLine[len(trimLine)-1]))
        sum = sum + tmpVal
    }
    fmt.Println(sum)
}
