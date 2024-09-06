/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var commands = map[string]interface{}{
	"add":   addFlashCard,
	"clear": clearScreen,
}

func addFlashCard() {
	fmt.Println("Add flash card")
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin": // darwin is for macOS
		cmd = exec.Command("clear")
	default:
		fmt.Println("Unsupported platform")
		return
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error clearing screen: %v\n", err)
	}
}

func printPrompt() {
	fmt.Print("clfashcards", "> ")
}

func parseInput(text string) []string {
	cleaned_text := strings.TrimSpace(strings.ToLower(text))
	text_list := strings.Fields(cleaned_text)
	return text_list
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clflashcards.git",
	Short: "A simple command line tool to create and review flahcards",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewScanner((os.Stdin))
		printPrompt()
		for reader.Scan() {
			// fmt.Println(commands)
			text_list := parseInput(reader.Text())
			if command, exists := commands[text_list[0]]; exists {
				// Call a hardcoded function
				command.(func())()
			} else if strings.EqualFold(".exit", text_list[0]) {
				// Close the program on the exit command
				return
			}
			printPrompt()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clflashcards.git.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
