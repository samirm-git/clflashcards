package flashcards

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

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

func removeStoreFromPath(path string) string {
	cleanPath := filepath.Clean(path)
	parts := strings.Split(cleanPath, string(filepath.Separator))

	if len(parts) > 0 && parts[0] == getStoreName() {
		shortened_path := filepath.Join(parts[1:]...)
		return shortened_path
	} else {
		return path
	}
}

func addTxtExtension(fname string) string {
	if !strings.HasSuffix(fname, ".txt") {
		return fname + ".txt"
	}
	return fname
}

func getLastDirAndFilename(path string) string {
	path = filepath.Clean(path)
	lastDir := filepath.Base(filepath.Dir(path))
	filename := filepath.Base(path)

	return filepath.Join(lastDir, filename)
}

func IntroPrint() {
	fmt.Println(`   ________       ________           __                        __
  / ____/ /      / ____/ /___ ______/ /_  _________ __________/ /____
 / /   / /      / /_  / / __ |/ ___/ __ \/ ___/ __ |/ ___/ __  / ___/
/ /___/ /___   / __/ / / /_/ (__  ) / / / /__/ /_/ / /  / /_/ (__  )
\____/_____/  /_/   /_/\__,_/____/_/ /_/\___/\__,_/_/   \__,_/____/   `)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Enter command or 'help' for the available commands.")
}

func printPromptCommand(command string) {
	colour := color.New(color.FgHiGreen).SprintFunc()
	fmt.Printf("  %s> ", colour(command))
}

func GetPrompt(currentCard string) string {
	if currentCard == "" {
		return "clflashcards> "
	} else {
		colour := color.New(color.FgCyan).SprintFunc()
		prompt := fmt.Sprintf("clflashcards %s> ", colour(getLastDirAndFilename(currentCard)))
		return prompt
	}
}

func getUserInput(currentCard string, scanner bufio.Scanner, command string) (string, error) {
	//currentCard not used. Remove later if this is still the case
	printPromptCommand(command)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("awaiting input:  %w", err)
	} else {
		input := scanner.Text()
		return input, nil
	}
}

func CheckQuit(text string) bool {
	if strings.ToLower(strings.TrimSpace(text)) == "quit" || strings.TrimSpace(strings.ToLower(text)) == "q" {
		return true
	}
	return false
}
