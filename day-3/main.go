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

	_, path := os.Args[1], os.Args[2]

	solve(path)
}

func solve(path string) {
	mem := parse(path)

	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	sum := 0

	for _, mul := range re.FindAllStringSubmatch(mem, -1) {
		nums := strings.Split(mul[0][4:len(mul[0])-1], ",")

		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])

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
