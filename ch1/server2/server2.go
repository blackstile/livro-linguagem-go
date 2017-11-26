// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Url digitada: %v\n", r.URL.Path)
	log.Printf("Valor count antes %d", count)
	mu.Lock()
	count++
	mu.Unlock()
	log.Printf("Valor count depois %d\n", count)
	fmt.Fprintf(w, "URL.Path =  %q\n", r.URL.Path)
}

//counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()

}
