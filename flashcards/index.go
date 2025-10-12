package flashcards

import (
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
