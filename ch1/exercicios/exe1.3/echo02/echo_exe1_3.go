package main

import (
	"fmt"
	"os"
	"time"
)

var s, sep string

func main() {
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "

	}
	fmt.Println(s)
	end := time.Now()
	fmt.Println("Duração For Range: ", end.Sub(start))
}
