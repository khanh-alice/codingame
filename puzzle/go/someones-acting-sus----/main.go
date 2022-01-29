package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Ship struct {
	rooms map[rune]*Room
}

func NewShip(len int, r string) Ship {
	rooms := make(map[rune]*Room, len)

	for i := 0; i < len; i++ {
		id := rune(r[i])
		rooms[id] = &Room{id: id}
	}

	for i := 0; i < len; i++ {
		room := rooms[rune(r[i])]

		if i == 0 {
			room.Prev = rooms[rune(r[len-1])]
		} else {
			room.Prev = rooms[rune(r[i-1])]
		}

		if i == len-1 {
			room.Next = rooms[rune(r[0])]
		} else {
			room.Next = rooms[rune(r[i+1])]
		}
	}

	return Ship{rooms: rooms}
}

func (s *Ship) Distance(from rune, to rune) int {
	if from == to {
		return 0
	}

	forward := 0
	for curr := from; curr != to; curr = s.rooms[curr].Next.id {
		forward++
	}

	backward := 0
	for curr := from; curr != to; curr = s.rooms[curr].Prev.id {
		backward++
	}

	return int(math.Min(float64(forward), float64(backward)))
}

func (s *Ship) IsSus(record string) bool {
	return s.isSus(record, 0)
}

func (s *Ship) isSus(record string, from int) bool {
	start := from
	for i := start; i < len(record); i++ {
		if record[i] != '#' {
			start = i
			break
		}
	}

	end := start + 1
	for i := end; i < len(record); i++ {
		if record[i] != '#' {
			end = i
			break
		}
	}

	if start >= len(record) || end >= len(record) {
		return false
	}

	if s.Distance(rune(record[start]), rune(record[end])) > end-start {
		return true
	}

	return s.isSus(record, end)
}

type Room struct {
	id   rune
	Prev *Room
	Next *Room
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L int
	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &L)

	scanner.Scan()
	F := scanner.Text()

	ship := NewShip(L, F)

	var N, K int
	scanner.Scan()
	_, _ = fmt.Sscan(scanner.Text(), &N, &K)

	for i := 0; i < N; i++ {
		scanner.Scan()
		record := scanner.Text()

		if ship.IsSus(record) {
			fmt.Println("SUS")
		} else {
			fmt.Println("NOT SUS")
		}
	}
}
