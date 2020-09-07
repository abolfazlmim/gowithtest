package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// usage go run fetch.go http://gopl.io
// usage go run fetch.go http://google.com

// bad request go run fetch.go http://bad.gopl.io
//fetch: Get http://bad.gopl.io: dial tcp: lookup bad.gopl.io on 103.86.99.100:53: no such host
//exit status 1
