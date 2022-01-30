package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	var M, N int
	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &M, &N)

	cols := make([]int, N)
	for i := 0; i < N; i++ {
		cols[i] = M
	}

	var T int
	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &T)

	choices := make([]int, T)
	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	for i := 0; i < T; i++ {
		c, _ := strconv.ParseInt(inputs[i], 10, 32)
		choices[i] = int(c)
	}

	for _, c := range choices {
		i := c - 1
		cols[i]--

		fmt.Println(calPerimeter(cols))
	}
}

func calPerimeter(cols []int) int {
	w := len(cols)
	for i := 0; i < len(cols); i++ {
		if cols[i] == 0 {
			w--
		}
	}
	w = w + w

	h := cols[0] + cols[len(cols)-1]
	for i := 1; i < len(cols); i++ {
		if cols[i] != cols[i-1] {
			h += int(math.Abs(float64(cols[i] - cols[i-1])))
		}
	}

	return w + h
}
