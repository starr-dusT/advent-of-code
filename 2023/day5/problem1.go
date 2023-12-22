package main

import (
	"fmt"
	L "github.com/starr-dusT/advent-of-code/lib"
)

func spaceToSlice(s string) []int {
    slice := []int{} 
    for _, strNumber := range strings.Split(s, " ") {
        slice = append(slice, L.StrToInt(strNumber))
    }
    return slice
}

func makeMap(ss []string) map[int](int) {
    return make(map[int](int))
}

func main() {
    seeds := []int{} 
    tmpAlmac := []string{}
    maps := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}
    input := []string{}
    for i, line := range L.Read_stdin() {
        if i == 0 {
            seeds = spaceToSlice(strings.Split(line, ": ")[1])
        } else if i > 2 && line != "" {
            fmt.Println(line)
            input = append(input, line)
        }
    }

    fmt.Println(seeds)
}
