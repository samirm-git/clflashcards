package main

import (

	// "flag"
	"fmt"
	"os"

	"github.com/samirm-git/clflashcards/flashcards"
)

var editorOptions = map[string]bool{"vim": true}

func main() {
	// config, err := loadConfig("config.json")
	// if err != nil {
	// 	fmt.Println("Error loading config:", err)
	// 	os.Exit(1)
	// }
	// Define command-line flags
	idx, err := flashcards.BuildFlashcardIndex()
	for dir, files := range idx.FilesByDir {
		fmt.Printf("Directory: %s\n", dir)
		fmt.Printf("  Files: %v\n", files)
		fmt.Println()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	flashcards.RunREPL(idx)

}
