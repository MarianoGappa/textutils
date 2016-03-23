package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func readLines(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines[scanner.Text()]++
	}
	return lines, scanner.Err()
}

func diff(lines1 map[string]int, lines2 map[string]int) {
	for k := range lines1 {
		if _, ok := lines2[k]; !ok {
			fmt.Println(k)
		}
	}
	return
}

func main() {
	flag.Parse()
	args := flag.Args()

	if args[0] == "diff" {
		lines1, _ := readLines(args[1])
		lines2, _ := readLines(args[2])

		diff(lines1, lines2)
	}
}
