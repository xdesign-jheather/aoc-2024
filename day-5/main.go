package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rules []Rule

func (r Rules) Before(page int, update Update) []int {
	var pages []int

	for _, rule := range r {
		if rule.Right != page {
			continue
		}

		if !update.Has(rule.Left) {
			continue
		}

		pages = append(pages, rule.Left)
	}

	return pages
}

type Rule struct {
	Left, Right int
}

type Updates []Update

type Update []int

func (u Update) Has(page int) bool {
	for _, p := range u {
		if p == page {
			return true
		}
	}
	return false
}

func (u Update) Middle() int {
	if len(u)%2 == 0 {
		log.Fatal()
	}

	i := len(u) - 1

	return u[i/2]
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
	rules, updates := parse(path)

	count := 0

	for _, update := range updates {
		if correct(update, rules) {
			fmt.Println(update)
			count += update.Middle()
		}
	}

	fmt.Println(count)
}

func correct(update Update, rules Rules) bool {
	seen := map[int]bool{}
	for _, page := range update {

		for _, other := range rules.Before(page, update) {
			if !seen[other] {
				return false
			}
		}

		seen[page] = true
	}
	return true
}

func parse(path string) (Rules, Updates) {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	var rules Rules
	var updates Updates

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.ContainsRune(line, '|'):
			parts := strings.Split(line, "|")
			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])
			rules = append(rules, Rule{
				Left:  left,
				Right: right,
			})
		case strings.ContainsRune(line, ','):
			var update Update

			for _, page := range strings.Split(line, ",") {
				num, err := strconv.Atoi(page)

				if err != nil {
					log.Fatal(err)
				}

				update = append(update, num)
			}

			updates = append(updates, update)
		}
	}

	return rules, updates
}
