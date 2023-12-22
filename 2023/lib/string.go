package lib

import (
	"strconv"
    "strings"
)

func StrToInt(str string) int {
  str = strings.TrimFunc(str, func(r rune) bool {
      return r < '0' || r > '9'
  })
  n, _ := strconv.Atoi(str)
  return n
}
