package lib

import (
	"strconv"
	"unicode"
)

func FindNumbers(inputStr string) ([]int, [][]int) {
    numbers := []int{}
    indexes := [][]int{}
    number := ""
    tempIndexes := []int{}
    for i, char := range inputStr {
        if number != "" && !unicode.IsNumber(char) {
            intNumber, _ := strconv.Atoi(number)
            numbers = append(numbers, intNumber)
            indexes = append(indexes, tempIndexes)
            number = ""
            tempIndexes = []int{}
        } else if unicode.IsNumber(char) {
            number = number + string(char)
            tempIndexes = append(tempIndexes, i)
        }
    }
    if number != "" {
        intNumber, _ := strconv.Atoi(number)
        numbers = append(numbers, intNumber)
        indexes = append(indexes, tempIndexes)
    }
    return numbers, indexes
}
