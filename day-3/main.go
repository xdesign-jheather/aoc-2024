package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	part, path := os.Args[1], os.Args[2]

	if part == "1" {
		solve1(path)
	}

	if part == "2" {
		solve2(path)
	}
}

func solve1(path string) {
	mem := parse(path)

	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	sum := 0

	for _, mul := range re.FindAllStringSubmatch(mem, -1) {
		n1, n2 := nums(mul[0])

		sum += n1 * n2
	}

	fmt.Println(sum)
}

func parse(path string) string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data, _ := io.ReadAll(f)

	return string(data)
}

func nums(input string) (int, int) {
	operands := strings.Split(input[4:len(input)-1], ",")

	n1, e1 := strconv.Atoi(operands[0])

	if e1 != nil {
		log.Fatal(e1)
	}

	n2, e2 := strconv.Atoi(operands[1])

	if e2 != nil {
		log.Fatal(e2)
	}

	return n1, n2
}

func solve2(path string) {
	mem := parse(path)

	re := regexp.MustCompile(`(mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\))`)

	enabled := true
	sum := 0

	for _, bit := range re.FindAllStringSubmatch(mem, -1) {

		switch {
		case bit[0] == "do()":
			enabled = true
			continue
		case bit[0] == "don't()":
			enabled = false
			continue
		case !enabled:
			continue
		}

		n1, n2 := nums(bit[0])

		sum += n1 * n2
	}

	fmt.Println(sum)
}
