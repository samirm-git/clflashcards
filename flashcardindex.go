package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FlashcardIndex struct {
	MasterStore string
	currentCard string
	FilesByDir  map[string][]string
}

func (idx *FlashcardIndex) SetCurrentFile(filepath string) {
	idx.currentCard = filepath
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

func BuildFlashcardIndex(masterDir string) (*FlashcardIndex, error) {
	idx := &FlashcardIndex{
		MasterStore: masterDir,
		currentCard: "",
		FilesByDir:  make(map[string][]string),
	}

	err := filepath.WalkDir(masterDir, idx.walkFunc)
	if err != nil {
		return nil, err
	}

	for dir, files := range idx.FilesByDir {
		fmt.Printf("Directory: %s\n", dir)
		fmt.Printf("  Files: %v\n", files)
		fmt.Println()
	}
	return idx, nil
}
