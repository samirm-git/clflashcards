package flashcards

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FlashcardIndex struct {
	MasterStore string
	CurrentCard string
	FilesByDir  map[string][]string
}

func (idx *FlashcardIndex) SetCurrentFile(filepath string) {
	idx.CurrentCard = filepath
}

func (idx *FlashcardIndex) walkFunc(path string, entry os.DirEntry, walkErr error) error {
	if walkErr != nil {
		return walkErr
	}

	if entry.IsDir() {
		return nil
	}

	directory := filepath.Dir(path)
	filename := entry.Name()

	idx.FilesByDir[directory] = append(idx.FilesByDir[directory], filename)
	return nil
}

func BuildFlashcardIndex(idx *FlashcardIndex) (*FlashcardIndex, error) {
	// If nil, create a new instance
	if idx == nil {
		idx = &FlashcardIndex{
			MasterStore: getStorePath(),
			CurrentCard: "",
			FilesByDir:  make(map[string][]string),
		}
		createIntroFile()
	} else {
		// Ensure FilesByDir is initialized
		if idx.FilesByDir == nil {
			idx.FilesByDir = make(map[string][]string)
		}

		for k := range idx.FilesByDir {
			delete(idx.FilesByDir, k)
		}
	}

	// Walk through the directory and update index
	err := filepath.WalkDir(idx.MasterStore, idx.walkFunc)
	if err != nil {
		return nil, err
	}

	return idx, nil
}

func getStoreName() string {
	return "clflashcards_home"
}

func getStorePath() string {
	home, err := os.UserHomeDir()
	check(err)
	store := filepath.Join(home, getStoreName())

	if _, err := os.Stat(store); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating flashcard home...")
		os.Mkdir(store, 0700)
	}
	return store
}

func createIntroFile() {
	home, err := os.UserHomeDir()
	check(err)
	introFile := filepath.Join(home, getStoreName(), "intro.txt")
	introText := []byte(`What is the purpose of this app? | To provide a simple and easy way to store ideas as flashcards. 
What are flascard files? | Simple txt files in which each line denotes a flashcard.
What are flaschards? | Question and Answer separated by a '|' operator.
Where are flashcard files stored? | All flashcard files are stored in your home directory in a dir called 'clflashcards_home'. DO NOT DELETE THIS.
How can I use this app? | Create txt files inside clflashcards_home. Add entries using 'create' or 'edit' command. Test yourself using 'testme'.
Extra | Flashcards can be created outside of the app using any txt editor. This cli app just allows for an easy way to do this direct from the command line.`)

	err = os.WriteFile(introFile, introText, 0644)
	check(err)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
