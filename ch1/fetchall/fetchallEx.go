//fetchallEx fetches urls sequentially and prints request times and size of body
//demonstrates that http.Get uses caching
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		fetch(url)
	}
}

func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetchall: %v\n", err)
		return
	}
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Printf("fetchall: %v\n", err)
		return
	}
	fmt.Printf("%.2fs\t%d7\n", time.Since(start).Seconds(), nbytes)
}
