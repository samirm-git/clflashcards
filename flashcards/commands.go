package flashcards

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AEROGU/tvchooser"
	"github.com/spf13/pflag"
)

func runSelectFlashCard(idx *FlashcardIndex, path string) error {
	path = addTxtExtension(path)
	abspath, err := idx.SmartCardFinder(path)

	if err != nil {
		return fmt.Errorf("selecting flashcard  %s  : %w", path, err)
	}

	idx.SetCurrentFile(abspath)
	fmt.Println("Successfully selected file.")
	return nil
}

func runSelectFlashCardGUI(idx *FlashcardIndex) error {
	clflashcards_home := getStorePath()
	idx.SetCurrentFile(tvchooser.FileChooser(nil, false, clflashcards_home))
	fmt.Println("Successfully selected file.")
	return nil
}

func saveFlashcard(savePath, question, answer string) error {
	// Open the file in append mode, creating it if it doesn't exist
	f, err := os.OpenFile(savePath, os.O_APPEND|os.O_WRONLY, 0644)
	if os.IsNotExist(err) {
		return fmt.Errorf("saving file %s  but file does not exist: %w", savePath, err)
	} else if err != nil {
		return fmt.Errorf("opening file %s : %w", savePath, err)
	}
	defer f.Close()

	// Write the flashcard entry
	entry := fmt.Sprintf("%s | %s\n", question, answer)
	_, err = f.WriteString(entry)
	if err != nil {
		return fmt.Errorf("saving %s : %w", savePath, err)
	} else {
		fmt.Println("Flashcard saved successfully")
	}
	return nil
}

func runShow(idx *FlashcardIndex, args []string) error {
	if idx.currentCard == "" {
		fmt.Println("NO FILE SELECTED. Use 'select' or 'edit' to change current flashcard.")
		return nil
	}
	content, err := os.ReadFile(idx.currentCard)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filepath.Base(idx.currentCard), err)
	}
	fmt.Println("Flashcards:\n", string(content))
	return nil
}

func runCreate(idx *FlashcardIndex, qaArgs []string) error {
	createfs := pflag.NewFlagSet("create", pflag.ContinueOnError)
	question := createfs.StringP("question", "q", "", "flashcard question")
	answer := createfs.StringP("answer", "a", "", "flashcard answer")

	if err := createfs.Parse(qaArgs); err != nil {
		return fmt.Errorf("parsing qaArgs: %w", err)
	}

	if len(qaArgs) < 2 {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Println("Enter a question or 'quit' to quit")
			questionScanner, err := getUserInput(idx.currentCard, *scanner)

			if err != nil {
				return fmt.Errorf("parsing question %s: %w", questionScanner, err)
			} else if checkQuit(questionScanner) {
				return nil
			}

			fmt.Println("Enter an answer or 'quit' to quit")
			answerScanner, err := getUserInput(idx.currentCard, *scanner)
			if err != nil {
				return fmt.Errorf("parsing answer %s: %w", answerScanner, err)
			} else if checkQuit(answerScanner) {
				return nil
			}

			err = saveFlashcard(idx.currentCard, questionScanner, answerScanner)
			if err != nil {
				return fmt.Errorf("saving question  %s  and answer  %q  : %w", question, answer, err)
			}
		}
	}

	if *question == "" || *answer == "" {
		*question = qaArgs[0]
		*answer = qaArgs[1]
	}

	err := saveFlashcard(idx.currentCard, *question, *answer)
	if err != nil {
		return fmt.Errorf("saving question  %s  and answer  %q  : %w", *question, *answer, err)
	}
	return nil
}

func runEditFile(idx *FlashcardIndex, args []string) error {
	editfs := pflag.NewFlagSet("editFile", pflag.ContinueOnError)
	fname := editfs.StringP("filename", "f", "", "Name of file to edit")
	editor := editfs.StringP("editor", "e", "code", "Text editor to open file with")
	isNewFile := editfs.BoolP("newFile", "n", false, "Flag to create new file in same dir as currently selected file, if fname does not exist")

	if err := editfs.Parse(args); err != nil {
		return fmt.Errorf("parsing args: %w", err)
	}

	if !editfs.Changed("filename") { //filename flag argument is not provided
		positionalArgs := editfs.Args()
		if len(positionalArgs) > 0 {
			*fname = positionalArgs[0] //check positional argument for filename
		} else {
			*fname = idx.currentCard
		}

		if *fname == "" {
			fmt.Println(`NO FILE SELECTED. Either use:
									1)  '-f' flag and specify file 
									or 2)	specify file without '-f' as first positional arg after 'edit' 
									or 3) use 'select' to change current flashcard.`)
			return nil
		}

	} else { //fname is specified
		*fname = addTxtExtension(*fname)
		fullcardpath, err := idx.SmartCardFinder(*fname)
		if err != nil {
			return err
		}
		idx.SetCurrentFile(fullcardpath)
		*fname = fullcardpath
	}

	if *isNewFile {
		f, err := os.OpenFile(*fname, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("opening file: %s: %w", *fname, err)
		}
		defer f.Close()
	} else {
		f, err := os.OpenFile(*fname, os.O_RDWR, 0644)
		if os.IsNotExist(err) {
			fmt.Printf("Error: File %s does not exist. If you want to create a new file with the edit command use -n flag\n", *fname)
			return nil
		} else if err != nil {
			return fmt.Errorf("opening file: %s: %w", *fname, err)
		}
		defer f.Close()
	}

	cmd := exec.Command(*editor, *fname)

	// Connect editor to terminal's stdin/stdout/stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run editor and wait until user exits
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("running editor: %s: %w", *editor, err)
	}
	fmt.Println("Editing finished!")

	idx.currentCard = *fname

	return nil

}
