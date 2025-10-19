package flashcards

import (
	"fmt"

	"github.com/fatih/color"
)

var helpCommandColour = color.New(color.FgBlue).SprintFunc()

func helpCreateFile() {
	fmt.Printf(` %s <path>
	RECOMMENDED TO USE NORMAL SHELL COMMANDS INSTEAD e.g. 'touch' or 'New-Item'. 

	Create new flashcardFile. <path> expects either absolute path from clflashcards_home or just the file name e.g. clflashcards_home/maths/unit1.txt or unit1.txt.
	If only file name is specified then the file will be created in currently selected file dir or in clflashcards_home if no card is currently selected. 
	If absolute path is given then the program will create any necessary dirs required.`, helpCommandColour("createFile"))
	fmt.Println()
}

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
	fmt.Printf(`  %s [n]
	Show n (or all) flashcards in the current file.`, helpCommandColour("show"))
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

func helpReresh() {
	fmt.Printf(`  %s
	Refresh flascard file index. Run this if you create or delete flashcard files.`, helpCommandColour("refresh"))
	fmt.Println()
}
