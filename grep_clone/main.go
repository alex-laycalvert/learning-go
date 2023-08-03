package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		panic("please provide a file and searchTerm\n")
	}
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("error reading file\n")
	}
	searchTerm := os.Args[2]
	data := split(string(bytes), '\n')
	matchingLines := make([]string, 0, len(data))
	if len(searchTerm) == 0 {
		matchingLines = data
	} else {
		for _, line := range data {
			for i := range line {
				for j, c := range searchTerm {
					if i+j > len(line)-1 {
						break
					}
					if string(c) != string(line[j+i]) {
						break
					}
					if j == len(searchTerm)-1 {
						matchingLines = append(matchingLines, line)
					}
				}
			}
		}
	}
	for _, l := range matchingLines {
		fmt.Println(l)
	}
}

func split(str string, del rune) []string {
	splits := make([]string, 0, len(str)+1)
	curr := ""
	for _, c := range str {
		if c == del {
			if len(curr) > 0 {
				splits = append(splits, curr)
			}
			curr = ""
			continue
		}
		curr += string(c)
	}
	if len(curr) > 0 {
		splits = append(splits, curr)
	}
	return splits
}
