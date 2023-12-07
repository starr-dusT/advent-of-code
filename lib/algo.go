package lib


func IndexIn(match [2]int, points [][2]int) bool {
    for _, point := range points {
        found := true
        for i := range point {
            if point[i] != match[i] {
                found = false 
            }
        }
        if found {
            return true
        }
    }
    return false
}

func In[T comparable](m T, s []T) bool {
    for _, v := range s {
        if v == m {
            return true
        }
    }
    return false
}

func Adjacent_indexes(blockIndexes [][2]int, capX int, capY int) [][2]int {
    indexes := [][2]int{}
    for _, blockIndex := range blockIndexes {
        for i := blockIndex[0] + 1; i > blockIndex[0] - 2; i-- {
            for j := blockIndex[1] + 1; j > blockIndex[1] - 2; j-- {
                newIndex := [2]int{i, j}
                if !IndexIn(newIndex, blockIndexes) && 
                   !IndexIn(newIndex, indexes) && 
                   !IndexIn(newIndex, [][2]int{blockIndex}) && 
                   i > -1 &&
                   j > -1 &&
                   i < capX &&
                   j < capY {
                    indexes = append(indexes, newIndex)
                }
            }
        }
    }
    return indexes
}
