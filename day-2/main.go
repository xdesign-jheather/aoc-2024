package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	part, path := os.Args[1], os.Args[2]

	if part == "1" {
		solve(path, false)
	}

	if part == "2" {
		solve(path, true)
	}
}

func parse(path string) [][]int {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var reports [][]int

	num := func(part string) int {
		v, err := strconv.Atoi(part)

		if err != nil {
			log.Fatal(err)
		}

		return v
	}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		var report []int

		for i := 0; i < len(parts); i++ {
			report = append(report, num(parts[i]))
		}

		reports = append(reports, report)
	}

	return reports
}

func solve(path string, dampener bool) {
	reports := parse(path)

	sum := 0

	for _, report := range reports {

		if safe(report) {
			sum += 1
			continue
		}

		if !dampener {
			continue
		}

		for i := 0; i < len(report); i++ {
			var sub []int

			for ii := 0; ii < len(report); ii++ {
				if ii == i {
					continue
				}

				sub = append(sub, report[ii])
			}

			if safe(sub) {
				sum += 1
				break
			}
		}
	}

	fmt.Println(sum)
}

func safe(report []int) bool {
	return reportSteady(report) && !reportVolatile(report)
}

func reportVolatile(report []int) bool {
	for i := 1; i < len(report); i++ {
		var diff int

		switch {
		case report[i] < report[i-1]: // Decreasing
			diff = report[i-1] - report[i]
		case report[i] > report[i-1]: // Increasing
			diff = report[i] - report[i-1]
		}

		if diff < 1 || diff > 3 {
			return true
		}
	}

	return false
}

func reportSteady(report []int) bool {
	inc, dec := false, false

	for i := 1; i < len(report); i++ {
		switch {
		case report[i] < report[i-1]: // Decreasing
			dec = true
		case report[i] > report[i-1]: // Increasing
			inc = true
		}

		if dec && inc {
			return false
		}
	}

	return true
}
