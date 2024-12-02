package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	path := os.Args[1]

	if strings.Contains(path, "1") {
		puzzle1(path)
	}
}

func parse(path string) [2][]int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	re := regexp.MustCompile(`^([0-9]+) +([0-9]+)$`)

	var lists [2][]int

	num := func(part string) int {
		v, err := strconv.Atoi(part)

		if err != nil {
			log.Fatal(err)
		}

		return v
	}

	for scanner.Scan() {
		line := scanner.Text()

		parts := re.FindStringSubmatch(line)

		if len(parts) != 3 {
			log.Fatalf("parts not %d length: %s", 3, line)
		}

		for i := 0; i < 2; i++ {
			lists[i] = append(lists[i], num(parts[i+1]))
		}
	}

	return lists
}

func puzzle1(path string) {
	// Parse the lists

	lists := parse(path)

	// Sort each list

	for i := range lists {
		sort.Ints(lists[i])
	}

	// Now do the math

	sum := 0

	for l := 0; l < len(lists[0]); l++ {
		switch {
		case lists[0][l] > lists[1][l]:
			sum += lists[0][l] - lists[1][l]
		default:
			sum += lists[1][l] - lists[0][l]
		}
	}

	fmt.Println(sum)
}
