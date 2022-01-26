package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	data []int
}

func NewStack(v int) Stack {
	return Stack{data: []int{v}}
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

type Terminal struct {
	stacks []Stack
}

func NewTerminal() Terminal {
	return Terminal{stacks: []Stack{}}
}

func (t *Terminal) Push(v int) {
	for i := 0; i < len(t.stacks); i++ {
		if t.stacks[i].Peek() >= v {
			t.stacks[i].Push(v)
			return
		}
	}

	t.stacks = append(t.stacks, NewStack(v))
}

func (t *Terminal) Len() int {
	return len(t.stacks)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		ans := solve(line)

		fmt.Println(ans)
	}
}

func solve(line string) int {
	t := NewTerminal()

	for _, char := range line {
		t.Push(int(char - 'A'))
	}

	return t.Len()
}
