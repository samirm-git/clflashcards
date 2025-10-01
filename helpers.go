package main

import (
	"bufio"
	"fmt"
	"os"
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

func getUserInput(scanner bufio.Scanner) (string, error) {
	printPrompt()
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("awaiting input:  %w", err)
	} else {
		input := scanner.Text()
		return input, nil
	}
}

func checkQuit(text string) bool {
	if strings.ToLower(strings.TrimSpace(text)) == "quit" || strings.TrimSpace(strings.ToLower(text)) == "q" {
		return true
	}
	return false
}

func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("checking if file exists: %w", err) // Return the actual error
	}
	return !info.IsDir(), nil
}

func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
