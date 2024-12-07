package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directions = []Letter{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

type Letter struct {
	Row, Col int
}

func (l Letter) Move(other Letter) Letter {
	return Letter{
		Row: l.Row + other.Row,
		Col: l.Col + other.Col,
	}
}

func (l Letter) Multiply(n int) Letter {
	return Letter{
		Row: l.Row * n,
		Col: l.Col * n,
	}
}

type Puzzle struct {
	chars []string
}

func (p Puzzle) MatchWord(letter Letter, direction Letter, word []byte) bool {
	for i := range word {
		next := letter.Move(direction.Multiply(i))

		if p.Letter(next) != word[i] {
			return false
		}
	}

	return true
}

func (p Puzzle) Letter(letter Letter) byte {
	if letter.Row < 0 || letter.Col < 0 {
		return '?'
	}

	if letter.Row > len(p.chars)-1 {
		return '?'
	}

	if letter.Col > len(p.chars[letter.Row])-1 {
		return '?'
	}

	return p.chars[letter.Row][letter.Col]
}

func (p Puzzle) Dimensions() (int, int) {
	if len(p.chars) == 0 {
		return 0, 0
	}

	return len(p.chars[0]), len(p.chars)
}

func (p Puzzle) Search(f func(Letter)) {
	width, height := p.Dimensions()

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			f(Letter{
				Row: h,
				Col: w,
			})
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	part, path := os.Args[1], os.Args[2]

	if part == "1" {
		solve1(path)
	}
}

func solve1(path string) {
	ws := parse(path)

	count := 0

	ws.Search(func(l Letter) {
		// Search in each direction

		for _, direction := range directions {
			if ws.MatchWord(l, direction, []byte("XMAS")) {
				count++
			}
		}
	})

	fmt.Println(count)
}

func parse(path string) *Puzzle {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	var result Puzzle

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		result.chars = append(result.chars, scanner.Text())
	}

	return &result
}
