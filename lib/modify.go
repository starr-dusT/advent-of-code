package lib

func MakeBlockFromLine(line int, indexes []int) [][2]int {
    block := [][2]int{}
    for _, index := range indexes {
        block = append(block, [2]int{index, line})
    }
    return block
}
