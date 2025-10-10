package main

import (
	"fmt"
	"os"

	"github.com/samirm-git/clflashcards/flashcards"
)

func main() {

	idx, err := flashcards.BuildFlashcardIndex()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	RunREPL(idx)

}
