//Dup2 prints the count and text of lines that appear more than once
//in the input. It reads from stdin or from a list of named files
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Store armazena os arquivos e quantidades
type Store struct {
	Files map[string]int
	Total int
}

func main() {
	counts := make(map[string]*Store)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for word, store := range counts {
		if store.Total > 1 {
			fmt.Printf("A Palavra %s ocorre:%d\r\n", strings.ToUpper(word), store.Total)
			for fileName, num := range store.Files {
				fmt.Printf("\tArquivo: %s tem um total de %d ocorrências\r\n", fileName, num)
			}
			fmt.Println()
		}
	}
}
func countLines(f *os.File, counts map[string]*Store) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		word := input.Text()
		store, found := counts[word]
		if !found {
			store = &Store{
				Files: make(map[string]int),
				Total: 0,
			}
		}
		store.Files[f.Name()]++
		store.Total++
		counts[word] = store
	}
	//Nota: ignorando erros em potencial de input.Err()
}
