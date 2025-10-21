package flashcards

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// runShellCommand executes a system shell command on Windows, macOS, or Linux.
// It respects the user's default shell when possible.
func RunShellCommand(idx *FlashcardIndex, args []string) error {
	if len(args) == 0 {
		return nil
	}

	if args[0] == "ls" {
		var currentDir string
		if idx.CurrentCard == "" {
			currentDir = idx.MasterStore
		} else {
			currentDir = filepath.Dir(idx.CurrentCard)
		}
		args = append([]string{"ls", currentDir}, args[1:]...)

	} else if args[0] == "cd" {
		return fmt.Errorf("cd command not supported. Please use select to select a flashcard file")
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		if isCommandAvailable("powershell") {
			cmd = exec.Command("powershell", "-NoLogo", "-NoProfile", "-Command", strings.Join(args, " "))
		} else {
			cmd = exec.Command("cmd", "/C", strings.Join(args, " "))
		}
	default:
		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "/bin/bash"
		}
		cmd = exec.Command(shell, "-c", strings.Join(args, " "))
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	return cmd.Run()
}

// isCommandAvailable checks if a given command exists in PATH
func isCommandAvailable(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
