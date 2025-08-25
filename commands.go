package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

func runCreate(question, answer string) {
	if question == "" || answer == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			var err error
			question, err = writeQ(*scanner)
			if err != nil || checkQuit(question) {
				break
			}

			answer, err = writeA(*scanner)
			if err != nil || checkQuit(answer) {
				break
			}

			saveFlashcard(question, answer)
		}
	} else {
		saveFlashcard(question, answer)
	}
}

func runEditFile(fname string) {
	fname = parsefname(fname)
	f, err := os.OpenFile(fname, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	f.Close()

	cmd := exec.Command("vim", fname)

	// Connect Vim to your terminal's stdin/stdout/stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run Vim and wait until user exits
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running vim:", err)
		return
	}

	fmt.Println("Editing finished!")

	if fname == "" {

	}
	// openInEditor(fname)
}
