package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chzyer/readline"
	"github.com/google/shlex"
	"github.com/samirm-git/clflashcards/flashcards"
)

func RunREPL2(idx *flashcards.FlashcardIndex) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          flashcards.GetPrompt(""),
		HistoryFile:     filepath.Join(os.TempDir(), "flashcards_history.tmp"), // saves command history
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	defer rl.Close()
	rl.CaptureExitSignal()

	flashcards.IntroPrint()
	for {
		rl.SetPrompt(flashcards.GetPrompt(idx.CurrentCard))
		line, err := rl.Readline()
		if err != nil { // io.EOF or readline.ErrInterrupt
			break
		}

		if flashcards.CheckQuit(line) {
			break
		}
		args, err := shlex.Split(line)
		if err != nil {
			fmt.Println("Unexpected Error parsing input :", err)
			continue
		}
		DispatchCommand(idx, args)
		fmt.Println()
		fmt.Printf("You entered: %s\n", line)
	}

}

func RunREPL(idx *flashcards.FlashcardIndex) {

	flashcards.IntroPrint()
	for {
		flashcards.PrintPrompt(idx.CurrentCard)
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
		if flashcards.CheckQuit(commandText) {
			break
		}
		args, err := shlex.Split(commandText)
		if err != nil {
			fmt.Println("Unexpected Error parsing input :", err)
			continue
		}
		DispatchCommand(idx, args)
		fmt.Println()
	}
}

func DispatchCommand(idx *flashcards.FlashcardIndex, args []string) {
	if len(args) == 0 {
		return
	}
	cmdName := args[0]
	cmd, ok := flashcards.GetCommand(cmdName)
	if !ok {
		fmt.Println("Unknown command:", cmdName)
		return
	}

	if err := cmd.Run(idx, args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
