package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/AEROGU/tvchooser"
)

const storename = "clflashcards_home"

var flashcard_path string

func getFlashcardsDir() string {
	home, _ := os.UserHomeDir()
	flashcard_dir := filepath.Join(home, storename)

	if _, err := os.Stat(flashcard_dir); errors.Is(err, os.ErrNotExist) {
		fmt.Println("===============================================================")
		fmt.Println("Creating flashcard home...")
		fmt.Println("===============================================================")
		os.Mkdir(flashcard_dir, 0700)
	}
	return flashcard_dir
}

func getShortenedFlashcardPath(path string) string {
	cleanPath := filepath.Clean(path)
	parts := strings.Split(cleanPath, string(filepath.Separator))

	if len(parts) > 0 && parts[0] == storename {
		shortened_path := filepath.Join(parts[1:]...)
		return shortened_path
	} else {
		return path
	}
}

func selectFlashCard(path string) error {
	shortened_path := getShortenedFlashcardPath(path)
	abspath := filepath.Join(getFlashcardsDir(), shortened_path)

	if _, err := os.Stat(abspath); err == nil {
		flashcard_path = abspath
		fmt.Println(flashcard_path)
	} else {
		return fmt.Errorf("error filenotfound %s: %w", abspath, err)
	}

	return nil
}

func selectFlashCardGUI() {
	clflashcards_home := getFlashcardsDir()
	flashcard_path = tvchooser.FileChooser(nil, false, clflashcards_home)
}

func getQ(scanner bufio.Scanner) (string, error) {
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

func getA(scanner bufio.Scanner) (string, error) {
	fmt.Println("Enter the answer or 'quit' to quit")
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
			question, err = getQ(*scanner)

			if err != nil {
				fmt.Println("Error parsing question input. Please try again.", err)
				continue

			} else if checkQuit(question) {
				break
			}

			answer, err = getA(*scanner)
			if err != nil {
				fmt.Println("Error parsing answer input. Please try again.", err)
				continue

			} else if checkQuit(answer) {
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
