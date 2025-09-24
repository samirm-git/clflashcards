package main

import (
	"bufio"
	// "flag"
	"fmt"
	"os"
	"strings"

	"github.com/google/shlex"
)

type Config struct {
	FlashcardDir string `json:"flashcard_dir"`
}

var editorOptions = map[string]bool{"vim": true}

func main() {
	// config, err := loadConfig("config.json")
	// if err != nil {
	// 	fmt.Println("Error loading config:", err)
	// 	os.Exit(1)
	// }
	// Define command-line flags
	if len(os.Args) > 1 {
		dispatchCommand(os.Args[1:])
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
		args, err := shlex.Split(commandText)
		if err != nil {
			fmt.Println("Unexpected Error parsing input :", err)
			continue
		}
		dispatchCommand(args)
	}
}

func dispatchCommand(args []string) {
	if len(args) == 0 {
		return
	}

	var err error
	err = nil
	switch args[0] {
	case "create":
		if len(args) < 3 {
			fmt.Println("Unexpected or missing arguments. Expected: create <question> <answer>")
			return
		}
		question := args[1]
		answer := args[2]
		err = runCreate(question, answer)

	case "editFile":
		if len(args) < 2 {
			fmt.Println("Unexpected of missing arguments. Expected: editFile <fname>")

		} else if len(args) == 3 {
			editorName := args[2]
			_, validEditor := editorOptions[editorName]
			if !validEditor {
				fmt.Printf("Invalid text editor option %s \n", editorName)
			}
			err = runEditFile(args[1], editorName)
		} else {
			err = runEditFile(args[1], "vim")
		}

	case "list":
		err = showFlashcards()

	case "select":
		if len(args) == 2 {
			fmt.Println("path inputted: ", args[1])
			err = selectFlashCard(args[1])
		} else {
			err = selectFlashCardGUI()
		}

	default:
		fmt.Println("Unknown command")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
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
