package flashcards

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/shlex"
)

type CommandFunc func(idx *FlashcardIndex, args []string) error

var commands = map[string]CommandFunc{
	"create":    func(idx *FlashcardIndex, args []string) error { return runCreate(idx, args) },
	"show":      func(idx *FlashcardIndex, args []string) error { return runShow(idx, args) },
	"edit":      func(idx *FlashcardIndex, args []string) error { return runEditFile(idx, args) },
	"editgui":   func(idx *FlashcardIndex, args []string) error { return runEditFileGUI(idx, args) },
	"selectgui": func(idx *FlashcardIndex, args []string) error { return runSelectFlashCardGUI(idx) },
	"select": func(idx *FlashcardIndex, args []string) error {
		if len(args) == 0 {
			return runSelectFlashCardGUI(idx)
		}
		return runSelectFlashCard(idx, args[0])
	},
}

func RunREPL(idx *FlashcardIndex) {

	introPrint()
	for {
		printPrompt(idx.currentCard)
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
		DispatchCommand(idx, args)
	}
}

func DispatchCommand(idx *FlashcardIndex, args []string) {
	if len(args) == 0 {
		return
	}
	cmdName := args[0]
	cmd, ok := commands[cmdName]
	if !ok {
		fmt.Println("Unknown command:", cmdName)
		return
	}

	if err := cmd(idx, args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
