package main

import (
	"fmt"
	"os"
	"strings"
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

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end = time.Now()

	fmt.Println("Duração Join: ", end.Sub(start))

	s, sep = "", ""
	start = time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	end = time.Now()
	fmt.Println(s)
	fmt.Println("Duração For i: ", end.Sub(start))
}
