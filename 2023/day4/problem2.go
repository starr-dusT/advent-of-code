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

func countCards(card int, allCards []string) []int {
    winCards := []int{}
    gameNumbers := strings.Split(strings.Split(allCards[card], ": ")[1], " | ")
    scratchNumbers := strings.Split(gameNumbers[1], " ")
    winningNumbers := strings.Split(gameNumbers[0], " ")
    for _, number := range scratchNumbers {
        if len(winCards) + card + 2 > len(allCards) {
            break
        }
        if L.In[string](number, winningNumbers) {
            winCards = append(winCards, card + len(winCards) + 1)
        }
    }
    return winCards
}

func countInst(m int, s []int) int {
    count := 0
    for _, val := range s {
        if m == val {
            count += 1
        }
    }
    return count
}

func main() {
    input := []string{}
    for _, line := range L.Read_stdin() {
        line = removeSpaces(line)
        input = append(input, line)
    }
    cards := []int{}
    dup := 0
    for i := range input {
        cards = append(cards, i)
        wonCards := countCards(i, input)
        for _, card := range wonCards {
            cards = append(cards, card)
        }
        for dup > 0 {
            for _, card := range wonCards {
                cards = append(cards, card)
            }
            dup -= 1
        }
        dup = countInst(i+1, cards)
    }
    fmt.Println(len(cards))
}
