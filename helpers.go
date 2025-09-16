package main

import (
	"fmt"
	"strings"
)

func parsefname(fname string) string {
	if strings.HasSuffix(fname, ".txt") {
		return fname
	}
	return fname + ".txt"
}

func introPrint() {
	fmt.Println(`   ________       ________           __                        __    
  / ____/ /      / ____/ /___ ______/ /_  _________ __________/ /____
 / /   / /      / /_  / / __ |/ ___/ __ \/ ___/ __ |/ ___/ __  / ___/
/ /___/ /___   / __/ / / /_/ (__  ) / / / /__/ /_/ / /  / /_/ (__  ) 
\____/_____/  /_/   /_/\__,_/____/_/ /_/\___/\__,_/_/   \__,_/____/   `)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Enter command or 'help' for the available commands.")
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
