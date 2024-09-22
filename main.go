package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	FlashcardDir string `json:"flashcard_dir"`
}

func main() {
	// config, err := loadConfig("config.json")
	// if err != nil {
	// 	fmt.Println("Error loading config:", err)
	// 	os.Exit(1)
	// }
	// Define command-line flags
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	question := createCmd.String("q", "", "Question for the flashcard")
	answer := createCmd.String("a", "", "Answer for the flashcard")

	editFileCmd := flag.NewFlagSet("editFile", flag.ExitOnError)
	subject_dir := createCmd.String("s", "", "Subject directory for flashcards")
	fname := flag.String("f", "", "Filename for flashcards")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Parse command-line arguments
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Expected 'create' or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		if *question == "" || *answer == "" {
			scanner := bufio.NewScanner(os.Stdin)
			for {
				var err error
				*question, err = getQ(*scanner)
				if err != nil {
					fmt.Println("Error reading question: ", err)
					break
				} else if checkQuit(*question) {
					break
				}

				*answer, err = getA(*scanner)
				if err != nil {
					fmt.Println("Error reading answer: ", err)
					break
				} else if checkQuit(*answer) {
					break
				}
				saveFlashcard(*question, *answer)
			}
		} else {
			fmt.Println("question ", *question)
			fmt.Println("answer ", *answer)
			saveFlashcard(*question, *answer)
		}
		os.Exit(1)
	case "editFile":
		editFileCmd.Parse(os.Args[2:])
		if *subject_dir == "" {
			scanner := bufio.NewScanner(os.Stdin)
			printPrompt()
			fmt.Println("Select subject or 'quit' to quit")
			//create functionality to list subject dirs and allow user to select one by pressing enter.
			if scanner.Scan() {
				subject_dir := scanner.Text()
				if checkQuit(subject_dir) {
					break
				}
			} else if err := scanner.Err(); err != nil {
				fmt.Println("Error reading subject: ", err)
			}

		}
		if *fname == "" {

		}

	case "list":
		listCmd.Parse(os.Args[2:])
		listFlashcards()
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}

// Load config from JSON file
// func loadConfig(path string) (*Config, error) {
// 	file, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var config Config
// 	if err := json.Unmarshal(file, &config); err != nil {
// 		return nil, err
// 	}
// 	return &config, nil
// }

//	func saveFlashcard(question, answer, flashcardDir string) {
//		f, err := os.OpenFile("testfile.txt", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
//		if err != nil {
//			fmt.Println("Error creating file: ", err)
//		}
//		defer f.Close()
//		entry := fmt.Sprintf("%s | %s\n", question, answer)
//		f.WriteString(entry)
//	}
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

func printPrompt() {
	fmt.Print("clflashcards> ")
}

func checkQuit(text string) bool {
	if strings.ToLower(strings.TrimSpace(text)) == "quit" || strings.TrimSpace(strings.ToLower(text)) == "q" {
		return true
	}
	return false
}
