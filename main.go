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
	// subject_dir := editFileCmd.String("s", "", "Subject directory for flashcards")
	fname := editFileCmd.String("f", "", "Filename for flashcards")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Parse command-line arguments
	if len(os.Args) > 2 {

		switch os.Args[1] {
		case "create":
			createCmd.Parse(os.Args[2:])
			fmt.Println("q", *question)
			fmt.Println("a", *answer)

			runCreate(*question, *answer)
			os.Exit(1)

		case "editFile":
			editFileCmd.Parse(os.Args[2:])
			runEditFile(*fname)

		case "list":
			listCmd.Parse(os.Args[2:])
			listFlashcards()

		default:
			fmt.Println("Unknown command")
			os.Exit(1)
		}
	} else {
		REPLMode()
	}
}

func REPLMode() {

	introPrint()
	for {
		printPrompt()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error getting command: ", err)
			break
		}

		commandText := strings.TrimSpace(scanner.Text())

		if commandText == "" {
			continue
		}
		if checkQuit(commandText) {
			break
		}
		parseCommand(commandText)
	}
}

func parseCommand(commandText string) {
	parts := strings.Fields(commandText)
	command := parts[0] //PARSE CREATE NOT WORKING CAUSE THIS JUST PARSES EACH WORD. NEED TO HANDLE REPL MODE LINE BY LNE
	args := parts[1:]

	switch command {
	case "create":
		runCreate(args[0], args[1])

	case "editFile":
		runEditFile(args[0])

	case "list":
		listFlashcards()

	default:
		fmt.Println("Unknown command")
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
