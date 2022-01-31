package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var W, H int
	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &W, &H)

	scanner.Scan()
	tops := strings.Split(scanner.Text(), "  ")

	d := make([]string, H-2)
	for i := 0; i < len(d); i++ {
		scanner.Scan()
		d[i] = scanner.Text()
	}

	scanner.Scan()
	bottoms := strings.Split(scanner.Text(), "  ")

	for i := len(d) - 1; i >= 0; i-- {
		row := d[i]

		for j := 0; j < len(bottoms)-1; j++ {
			if row[j*3+1] == '-' {
				bottoms[j], bottoms[j+1] = bottoms[j+1], bottoms[j]
			}
		}
	}

	for i := 0; i < len(tops); i++ {
		fmt.Printf("%s%s\n", tops[i], bottoms[i])
	}
}
