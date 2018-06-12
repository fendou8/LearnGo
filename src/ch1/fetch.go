package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reading%s:%v\n", url, err)
		}
		fmt.Printf("%s\n", b)
		status := resp.Status

		fmt.Printf("%s\n", status)
	}
}
