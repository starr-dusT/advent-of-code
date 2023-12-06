package main

import (
	"fmt"
    "strings"
	L "github.com/starr-dusT/advent-of-code/lib"
)

func removeSpaces(s string) string {
    lastChar := "0"
    rString := ""
    for i := range s {
        char := string(s[i])
        if char != " " || (char == " " && lastChar != " ") {
            rString = rString + char
        }
        lastChar = char
    }
    return rString
}

func main() {
    ans := 0
    for _, line := range L.Read_stdin() {
        line = removeSpaces(line)
        gameNumbers := strings.Split(strings.Split(line, ": ")[1], " | ")
        scratchNumbers := strings.Split(gameNumbers[1], " ")
        winningNumbers := strings.Split(gameNumbers[0], " ")
        score := 0
        for _, number := range scratchNumbers {
            if L.In[string](number, winningNumbers) {
                if score == 0 {
                    score += 1
                } else {
                    score *= 2
                }
            }
        }
        ans += score
    }
    fmt.Println(ans)
}
