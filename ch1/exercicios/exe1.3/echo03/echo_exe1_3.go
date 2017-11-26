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
	fmt.Println(strings.Join(os.Args[1:], " "))
	end := time.Now()
	fmt.Println("Duração Join: ", end.Sub(start))

}
