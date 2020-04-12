//Dup2 two prints the count and text of duplicate lines. It read lines from a list if files
//or from stdin
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	occursIn := make(map[string]map[string]int)
	if len(os.Args[1:]) == 0 {
		countLines(os.Stdin, counts, occursIn)
	} else {
		for _, filename := range os.Args[1:] {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occursIn)
			f.Close()
		}
	}
	for key, value := range counts {
		if value > 1 {
			sep := ""
			var filesStr string
			for fileName, occurences := range occursIn[key] {
				filesStr = fmt.Sprintf("%s%s%s (%d)", filesStr, sep, fileName, occurences)
				sep = " "
			}
			fmt.Printf("%d\t%s\t%s\n", value, key, filesStr)
		}
	}
}

func countLines(file *os.File, counts map[string]int,
	occursIn map[string]map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
		if occursIn[input.Text()] == nil {
			occursIn[input.Text()] = make(map[string]int)
		}
		occursIn[input.Text()][file.Name()]++
	}
}
