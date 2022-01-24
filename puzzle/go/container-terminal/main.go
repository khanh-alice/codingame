package main

import "fmt"
import "os"
import "bufio"

type Stack struct {
	data []int
}

func NewStack() Stack {
	return Stack{data: []int{}}
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Peek() int {
	l := len(s.data)
	return s.data[l-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

type Terminal struct {
	stacks []Stack
}

const TerminalSize = 27

func NewTerminal() Terminal {
	stacks := make([]Stack, TerminalSize)

	for i := 0; i < TerminalSize; i++ {
		stacks[i] = NewStack()
	}

	return Terminal{stacks: stacks}
}

func (t *Terminal) Push(v int) {
	if !t.stacks[v].IsEmpty() {
		t.stacks[v].Push(v)
		return
	}

	for i := v + 1; i < TerminalSize; i++ {
		if !t.stacks[i].IsEmpty() {
			t.stacks[i].Push(v)
			t.stacks[v] = t.stacks[i]
			t.stacks[i] = NewStack()
			return
		}
	}

	t.stacks[v].Push(v)
}

func (t *Terminal) Count() int {
	c := 0
	for i := 0; i < TerminalSize; i++ {
		if !t.stacks[i].IsEmpty() {
			c++
		}
	}
	return c
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

	return t.Count()
}
