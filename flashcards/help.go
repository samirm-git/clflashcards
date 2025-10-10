package flashcards

import (
	"fmt"

	"github.com/fatih/color"
)

var helpCommandColour = color.New(color.FgBlue).SprintFunc()

func helpCreate() {
	fmt.Printf(`  %s [--question] [--answer]
	Create new flashcard in currently selected flashcard file`, helpCommandColour("create"))
	fmt.Println()
}

func helpEdit() {
	fmt.Printf(`  %s [--filename] [--editor] [--isNewFile]
	Edit selected file (either currently selected or --filename)`, helpCommandColour("edit"))
	fmt.Println()
}

func helpEditGui() {
	fmt.Printf(`  %s
	Open a GUI to edit a flashcard file.`, helpCommandColour("editgui"))
	fmt.Println()
}

func helpSelect() {
	fmt.Printf(`  %s <path>
	Select new flashcard file. <path> can be relative to clflashcards_store or current flashcard dir.`, helpCommandColour("select"))
	fmt.Println()
}

func helpSelectGui() {
	fmt.Printf(`  %s
	Open a GUI to selecy a flashcard file.`, helpCommandColour("selectGui"))
	fmt.Println()
}

func helpShow() {
	fmt.Printf(`  %s
	STILL IN PROGRESS.`, helpCommandColour("show"))
	fmt.Println()
}

func helpTestMe() {
	fmt.Printf(`  %s [--random] [--numberOfQuestions]
	Test your flashcard (currently selected flashcard) understanding by entering the answers to the questions).`, helpCommandColour("testme"))
	fmt.Println()
}

func helpHelp() {
	fmt.Printf(`  %s [command]
	Show help option. Optional argument for help on a specific [command]`, helpCommandColour("help"))
	fmt.Println()
}
