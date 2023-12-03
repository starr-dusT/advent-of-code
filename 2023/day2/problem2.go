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
    sum := 0
    for _, line := range read_stdin() {
        mostRed := 1
        mostGreen := 1
        mostBlue := 1
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
                        if mostRed < last {
                            mostRed = last
                        }
                    case "blue":
                        if mostBlue < last {
                            mostBlue = last 
                        }
                    case "green":
                        if mostGreen < last {
                            mostGreen = last 
                        }
                    default:
                }
            }
        }
        sum = sum + mostBlue*mostGreen*mostRed
    }
    fmt.Println(sum)
}
