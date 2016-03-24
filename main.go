package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func readLines(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineWords := strings.Fields(line)
		for _, v := range lineWords {
			if filtered := filterChars(v); len(filtered) > 0 {
				words[filtered]++
			}
		}
	}
	return words, scanner.Err()
}

func filterChars(s string) string {
	result := ""
	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') ||
			c == 'á' || c == 'é' || c == 'í' || c == 'ó' || c == 'ú' ||
			c == 'Á' || c == 'É' || c == 'Í' || c == 'Ó' || c == 'Ú' ||
			c == 'ñ' || c == 'Ñ' {

			switch {
			case c >= 'A' && c <= 'Z':
				c += 'a' - 'A'
			case c == 'Á':
				c = 'á'
			case c == 'É':
				c = 'é'
			case c == 'Í':
				c = 'í'
			case c == 'Ó':
				c = 'ó'
			case c == 'Ú':
				c = 'ú'
			case c == 'Ñ':
				c = 'ñ'
			}

			result = result + string(c)
		}
	}
	return result
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

	switch args[0] {
	case "freq":
		lines, _ := readLines(args[1])
		for k, v := range lines {
			fmt.Println(k, v)
		}
	case "diff":
		lines1, _ := readLines(args[1])
		lines2, _ := readLines(args[2])

		diff(lines1, lines2)
	}

}
