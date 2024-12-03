package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	for day := 1; day <= 25; day++ {
		run(day)
	}
}

func run(day int) {
	path := filepath.Join("..", "day-"+strconv.Itoa(day))

	stat, err := os.Stat(path)

	if err != nil || !stat.IsDir() {
		return
	}

	if err = build(path); err != nil {
		return
	}

	if err = part(path, "1"); err != nil {
		return
	}

	if err = part(path, "2"); err != nil {
		return
	}
}

func part(path string, part string) error {
	start := time.Now()

	cmd := exec.Command("./aoc", part, "input.txt")
	cmd.Dir = path
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	fmt.Printf("Completed %s part %s in %s\n", filepath.Base(path), part, time.Since(start).Truncate(time.Millisecond))

	return err
}

func build(path string) error {
	cmd := exec.Command("go", "build", "-o", "aoc", ".")
	cmd.Dir = path
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
