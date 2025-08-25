package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeQ(scanner bufio.Scanner) (string, error) {
	fmt.Println("Enter a question or 'quit' to quit")
	printPrompt()
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	} else {
		newquestion := scanner.Text()
		return newquestion, nil
	}
}

func writeA(scanner bufio.Scanner) (string, error) {
	fmt.Println("Enter the enswer or 'quit' to quit")
	printPrompt()
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	} else {
		newasnwer := scanner.Text()
		return newasnwer, nil
	}
}

func saveFlashcard(question, answer string) {
	// Open the file in append mode, creating it if it doesn't exist
	f, err := os.OpenFile("testfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	// Write the flashcard entry
	entry := fmt.Sprintf("%s | %s\n", question, answer)
	_, err = f.WriteString(entry)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Flashcard saved successfully to", "testfile.txt")
	}
}

func listFlashcards() {
	filename := "flashcards.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading flashcards:", err)
		return
	}
	fmt.Println("Flashcards:\n", string(content))
}
