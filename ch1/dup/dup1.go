//prints duplicate lines on standard input as well as the number of times the lines appears
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	} //on make, ctrl-d for EOF

	for key, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, key)
		}
	}
}
