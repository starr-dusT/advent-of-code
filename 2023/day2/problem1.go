package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
    red := 12
    green := 13
    blue := 14
    sum := 0
    for j, line := range read_stdin() {
        possible := true
        line = strings.Replace(line,",", "", -1)
        gameStr := strings.Split(line, ": ")[1]
        for _, setStr := range strings.Split(gameStr, "; ") {
            chunks := strings.Split(setStr, " ")
            last := 0
            for i, chunk := range chunks {
                if i > 0 {
                    last, _ = strconv.Atoi(chunks[i-1])
                }
                switch chunk {
                    case "red":
                        if red < last {
                            possible = false
                        }
                    case "blue":
                        if blue < last {
                            possible = false
                        }
                    case "green":
                        if green < last {
                            possible = false
                        }
                    default:
                }
            }
        }
        if possible {
            sum = sum + j + 1
        }
    }
    fmt.Println(sum)
}
