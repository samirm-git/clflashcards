package main

import (
	"fmt"
	"strings"
)

func parsefname(fname string) string {
	if strings.HasSuffix(fname, ".txt") {
		return fname
	}
	return fname + ".txt"
}

func printPrompt() {
	fmt.Print("clflashcards> ")
}

func checkQuit(text string) bool {
	if strings.ToLower(strings.TrimSpace(text)) == "quit" || strings.TrimSpace(strings.ToLower(text)) == "q" {
		return true
	}
	return false
}
