package main

import (
	"bufio"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func addTxtExtension(fname string) string {
	if !strings.HasSuffix(fname, ".txt") {
		return fname + ".txt"
	}
	return fname
}

func introPrint() {
	fmt.Println(`   ________       ________           __                        __    
  / ____/ /      / ____/ /___ ______/ /_  _________ __________/ /____
 / /   / /      / /_  / / __ |/ ___/ __ \/ ___/ __ |/ ___/ __  / ___/
/ /___/ /___   / __/ / / /_/ (__  ) / / / /__/ /_/ / /  / /_/ (__  ) 
\____/_____/  /_/   /_/\__,_/____/_/ /_/\___/\__,_/_/   \__,_/____/   `)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Enter command or 'help' for the available commands.")
}

func printPrompt() {
	if currFlashCardPath == "" {
		fmt.Print("clflashcards> ")
	} else {
		colour := color.New(color.FgCyan).SprintFunc()
		fmt.Printf("clflashcards %s> ", colour(filepath.Base(currFlashCardPath)))
	}
}

func getQ(scanner bufio.Scanner) (string, error) {
	fmt.Println("Enter a question or 'quit' to quit")
	printPrompt()
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("awaiting question:  %w", err)
	} else {
		newquestion := scanner.Text()
		return newquestion, nil
	}
}

func getA(scanner bufio.Scanner) (string, error) {
	fmt.Println("Enter the answer or 'quit' to quit")
	printPrompt()
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("awaiting answer: %w", err)
	} else {
		newasnwer := scanner.Text()
		return newasnwer, nil
	}
}

func checkQuit(text string) bool {
	if strings.ToLower(strings.TrimSpace(text)) == "quit" || strings.TrimSpace(strings.ToLower(text)) == "q" {
		return true
	}
	return false
}
