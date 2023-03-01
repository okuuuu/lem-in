package main

import (
	"bufio"
	"main/antfarm"
	"os"
)

func main() {
	// path := os.Args[1]
	path := "./example00.txt"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	err = antfarm.Solve(scanner)
	if err != nil {
		panic(err)
	}
}
