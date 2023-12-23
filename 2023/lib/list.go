package lib

import (
)

func cut(i int, xs []int) (int, []int) {
  y := xs[i]
  ys := append(xs[:i], xs[i+1:]...)
  return y, ys
}
