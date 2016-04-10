package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func readWords(path string) (map[string]int, error) {
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

func readDict(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		section := strings.Split(line, "===")
		lineWords := strings.Split(section[1], ";")
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

func diff(words1 map[string]int, words2 map[string]int) {
	for k := range words1 {
		if _, ok := words2[k]; !ok {
			fmt.Println(k)
		}
	}
	return
}

func intersect(words1 map[string]int, words2 map[string]int) {
	for k := range words1 {
		if _, ok := words2[k]; ok {
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
		lines, _ := readWords(args[1])
		for k, v := range lines {
			fmt.Println(k, v)
		}
	case "diff":
		words1, _ := readWords(args[1])
		words2, _ := readWords(args[2])

		diff(words1, words2)
	case "adjectives":
		words1, _ := readWords(args[1])
		words2, _ := readDict("dictionaries/adjDic.txt")

		intersect(words1, words2)
	case "nouns":
		words1, _ := readWords(args[1])
		words2, _ := readDict("dictionaries/nounDic.txt")

		intersect(words1, words2)
	case "verbs":
		words1, _ := readWords(args[1])
		words2, _ := readDict("dictionaries/verbDic.txt")

		intersect(words1, words2)
	}

}
