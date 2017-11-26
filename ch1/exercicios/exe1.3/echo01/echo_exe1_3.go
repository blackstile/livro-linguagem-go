package main

import (
	"fmt"
	"os"
	"time"
)

var s, sep string

func main() {
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	end := time.Now()
	fmt.Println(s)
	fmt.Println("Duração For i: ", end.Sub(start))
}
