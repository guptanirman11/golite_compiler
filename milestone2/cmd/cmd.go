package main

import (
	"fmt"
	"milestone2/scanner"
	"milestone2/token"
	"os"
)

func main() {
	file_name := os.Args[1]
	dat, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	input := string(dat)
	s := scanner.New(input)
	var t token.Token

	for {
		t = s.NextToken()
		if t.Type == "EOF" {
			break
		} else if t.Type == "INVALID" {
			// An invalid token was found
			fmt.Println("ERROR: Illegal Token")
			break
		}
		t.Print()
	}
}
