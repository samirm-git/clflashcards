package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
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
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Expected 'create' or 'editFile' or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		fmt.Println("q", *question)
		fmt.Println("a", *answer)
		if *question == "" || *answer == "" {
			scanner := bufio.NewScanner(os.Stdin)
			for {
				var err error
				*question, err = writeQ(*scanner)
				if err != nil {
					fmt.Println("Error reading question: ", err)
					break
				} else if checkQuit(*question) {
					break
				}

				*answer, err = writeA(*scanner)
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
		*fname = parsefname(*fname)
		fmt.Println("fname", *fname)
		f, err := os.OpenFile(*fname, os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		f.Close()

		cmd := exec.Command("vim", *fname)

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
