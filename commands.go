package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/AEROGU/tvchooser"
	"github.com/spf13/pflag"
)

const storename = "clflashcards_home"

var currFlashCardPath string

func getFlashcardsDir() string {
	home, _ := os.UserHomeDir()
	flashcard_dir := filepath.Join(home, storename)

	if _, err := os.Stat(flashcard_dir); errors.Is(err, os.ErrNotExist) {
		fmt.Println("===============================================================")
		fmt.Println("Creating flashcard home...")
		fmt.Println("===============================================================")
		os.Mkdir(flashcard_dir, 0700)
	}
	return flashcard_dir
}

func removeStoreFromPath(path string) string {
	cleanPath := filepath.Clean(path)
	parts := strings.Split(cleanPath, string(filepath.Separator))

	if len(parts) > 0 && parts[0] == storename {
		shortened_path := filepath.Join(parts[1:]...)
		return shortened_path
	} else {
		return path
	}
}

func selectFlashCard(path string) error {
	shortened_path := addTxtExtension(removeStoreFromPath(path))
	abspath := filepath.Join(getFlashcardsDir(), shortened_path)

	if _, err := os.Stat(abspath); err == nil {
		currFlashCardPath = abspath
		fmt.Println(currFlashCardPath)
	} else {
		return fmt.Errorf("filenotfound %s: %w", abspath, err)
	}

	fmt.Println("Successfully selected file.")
	return nil
}

func selectFlashCardGUI() error {
	clflashcards_home := getFlashcardsDir()
	currFlashCardPath = tvchooser.FileChooser(nil, false, clflashcards_home)
	fmt.Println("Successfully selected file.")
	return nil
}

func saveFlashcard(question, answer string) error {
	// Open the file in append mode, creating it if it doesn't exist
	fname := "testfile.txt"
	f, err := os.OpenFile("testfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("opening file %s : %w", fname, err)
	}
	defer f.Close()

	// Write the flashcard entry
	entry := fmt.Sprintf("%s | %s\n", question, answer)
	_, err = f.WriteString(entry)
	if err != nil {
		return fmt.Errorf("saving %s : %w", fname, err)
	} else {
		fmt.Println("Flashcard saved successfully to", "testfile.txt")
	}
	return nil
}

func showFlashcards() error {
	f := "flashcards.txt"
	content, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", f, err)
	}
	fmt.Println("Flashcards:\n", string(content))
	return nil
}

func runCreate(question, answer string) error {
	var err error
	if question == "" || answer == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			question, err = getQ(*scanner)

			if err != nil {
				return fmt.Errorf("parsing question %s: %w", question, err)
			} else if checkQuit(question) {
				break
			}

			answer, err = getA(*scanner)
			if err != nil {
				return fmt.Errorf("parsing answer %s: %w", answer, err)
			} else if checkQuit(answer) {
				break
			}

			err = saveFlashcard(question, answer)
			if err != nil {
				return fmt.Errorf("saving question  %s  and answer  %q  : %w", question, answer, err)
			}
		}
	} else {
		err = saveFlashcard(question, answer)
		if err != nil {
			return fmt.Errorf("saving question  %s  and answer  %q  : %w", question, answer, err)
		}
	}
	return nil
}

func runEditFile(args []string) error {
	editfs := pflag.NewFlagSet("editFile", pflag.ContinueOnError)
	fname := editfs.StringP("filename", "f", "", "Name of file to edit")
	editor := editfs.StringP("editor", "e", "code", "Text editor to open file with")
	isNewFile := editfs.BoolP("newFile", "n", false, "Flag to create new file in same dir as currently selected file, if fname does not exist")

	if err := editfs.Parse(args); err != nil {
		return fmt.Errorf("parsing args: %w", err)
	}

	if !editfs.Changed("filename") {
		if currFlashCardPath == "" {
			fmt.Println("No file selected. Either use 'select' command or add file name as argument to 'edit'")
			return nil
		} else {
			*fname = currFlashCardPath
		}
	} else { //fname is specified
		*fname = addTxtExtension(*fname)
		if currFlashCardPath == "" { //not selected any flashcard
			*fname = filepath.Join(getFlashcardsDir(), *fname)
		} else {
			*fname = filepath.Join(filepath.Dir(currFlashCardPath), *fname)
		}
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

	currFlashCardPath = *fname

	return nil

	//Show what flashcard is currently selected in the printprompt()
}
