package lib

import (
	"bufio"
	"fmt"
	"os"
)

func Read_stdin() []string {
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
