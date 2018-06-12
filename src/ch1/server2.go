package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/hander", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method,r.URL,r.Proto)
	for k,v := range r.Header  {
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k,v := range r.Form  {
		fmt.Fprintf(w,"Form[%q] = %q\n",k,v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
