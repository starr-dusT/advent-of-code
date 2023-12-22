package lib

import (
	"strconv"
    "strings"
)

func cut(i int, xs []int) (int, []int) {
  y := xs[i]
  ys := append(xs[:i], xs[i+1:]...)
  return y, ys
}
