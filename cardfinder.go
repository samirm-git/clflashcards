package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (idx *FlashcardIndex) SmartCardFinder(inputPath string) (string, error) {
	var err error
	masterStore, err := filepath.Abs(idx.MasterStore)
	if err != nil {
		return "", err
	}
	currentCard, err := filepath.Abs(idx.currentCard)
	if err != nil {
		return "", err
	}
	activeDir := filepath.Dir(currentCard)
	inputPath = filepath.Clean(inputPath)

	// 1) Absolute path
	if filepath.IsAbs(inputPath) {
		isFileExists, err := fileExists(inputPath)
		if err != nil {
			return "", err
		}
		if isFileExists {
			return inputPath, nil
		}

	} else if strings.Contains(inputPath, string(os.PathSeparator)) {
		// 2) inputPath contains path separators -> try ancestor-aware resolution from active dir back to its parents

		if isSubpath(masterStore, activeDir) {
			cardFinder := activeDir
			parent := filepath.Dir(cardFinder)

			for !samePath(cardFinder, masterStore) && parent != cardFinder {
				candidatePath := filepath.Clean(filepath.Join(cardFinder, inputPath))
				candidateDirPath := filepath.Dir(candidatePath)

				if checkCardInIndex(idx, candidateDirPath) {
					return candidatePath, nil
				}
				parent = filepath.Dir(cardFinder)
				if parent == cardFinder {
					break
				}
				cardFinder = parent
			}

			candidatePath := filepath.Clean(filepath.Join(masterStore, inputPath))
			candidateDirPath := filepath.Dir(candidatePath)
			if checkCardInIndex(idx, candidateDirPath) {
				return candidatePath, nil
			}
			return "", fmt.Errorf("file not found: %s (tried active path ancestors and masterStore)", inputPath)
		}

	} else {
		fmt.Println("Got to else statement")
		fmt.Println(activeDir)
		// 3) No path separator -> prefer active directory first
		inputFile := inputPath
		fmt.Println(inputFile)
		fmt.Println("")
		if checkCardInIndex(idx, activeDir) {
			return filepath.Join(activeDir, inputFile), nil
		}
	}

	return "", fmt.Errorf("could not find input in flashcard store : %s", inputPath)
}

func checkCardInIndex(idx *FlashcardIndex, cardDirPath string) bool {
	//ONLY CHECKING IF THE DIR PATH GIVEN IN INPUT (SMART ADJUSTED) EXSITS IN INDEX
	//NOT CHECKING IF THE LAST PART I.E. FILE NAME MATCHES WITH A VALUE IN THE INDEX[DIR_PATH]
	if _, ok := idx.FilesByDir[cardDirPath]; ok {
		return true
	} else {
		return false
	}
}

func isSubpath(base, target string) bool {
	base = filepath.Clean(base)
	target = filepath.Clean(target)
	rel, err := filepath.Rel(base, target)
	if err != nil {
		return false
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(os.PathSeparator))
}

func samePath(a, b string) bool {
	a = filepath.Clean(a)
	b = filepath.Clean(b)
	return a == b
}
