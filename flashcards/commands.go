package flashcards

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/AEROGU/tvchooser"
	"github.com/spf13/pflag"
)

type Flashcard struct {
	Question string
	Answer   string
}

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
			questionScanner, err := getUserInput(idx.currentCard, *scanner, "question")

			if err != nil {
				return fmt.Errorf("parsing question %s: %w", questionScanner, err)
			} else if checkQuit(questionScanner) {
				return nil
			}

			fmt.Println("Enter an answer or 'quit' to quit")
			answerScanner, err := getUserInput(idx.currentCard, *scanner, "answer")
			if err != nil {
				return fmt.Errorf("parsing answer %s: %w", answerScanner, err)
			} else if checkQuit(answerScanner) {
				return nil
			}

			err = saveFlashcard(idx.currentCard, questionScanner, answerScanner)
			if err != nil {
				return fmt.Errorf("saving question  %s  and answer  %q  : %w", *question, *answer, err)
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

func runTestme(idx *FlashcardIndex, args []string) error {
	testmefs := pflag.NewFlagSet("testme", pflag.ContinueOnError)
	isRandomOrder := testmefs.BoolP("random", "r", false, "Flag to set question order to random")
	nquestions := testmefs.IntP("numberOfQuestion", "n", -1, "Number of questions to test")

	if err := testmefs.Parse(args); err != nil {
		return fmt.Errorf("parsing args: %w", err)
	}

	if idx.currentCard == "" {
		return fmt.Errorf("no card currently selected")
	}

	cards, err := parseFlashcardFile(idx.currentCard)
	if err != nil {
		return err
	}

	if *isRandomOrder {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	}

	if *nquestions <= 0 || *nquestions > len(cards) {
		*nquestions = len(cards)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i, card := range cards[:*nquestions] {
		fmt.Printf("\nQuestion %d/%d: %s\n", i+1, *nquestions, card.Question)
		fmt.Println("Your answer (or type 'quit' to exit): ")

		answerInput, err := getUserInput(idx.currentCard, *scanner, "answer")
		if err != nil {
			return fmt.Errorf("parsing answer input %s: %w", answerInput, err)
		} else if checkQuit(answerInput) {
			return nil
		}
		fmt.Println("Correct Answer:")
		fmt.Printf(" %s \n", card.Answer)
	}

	return nil
}

func parseFlashcardFile(fname string) ([]Flashcard, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", fname, err)
	}
	defer f.Close()
	var cards []Flashcard
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			continue // skip malformed lines
		}
		question := strings.TrimSpace(parts[0])
		answer := strings.TrimSpace(parts[1])
		cards = append(cards, Flashcard{Question: question, Answer: answer})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading file %s: %w", fname, err)
	} else {
		return cards, nil
	}
}

func runEditFileGUI(idx *FlashcardIndex, args []string) error {
	editguifs := pflag.NewFlagSet("editFileGUI", pflag.ContinueOnError)
	editor := editguifs.StringP("editor", "e", "code", "Text editor to open file with")
	isNewFile := editguifs.BoolP("newFile", "n", false, "Flag to create new file in same dir as currently selected file if fname does not exist.")

	if err := editguifs.Parse(args); err != nil {
		return fmt.Errorf("parsing args: %w", err)
	}
	clflashcards_home := getStorePath()
	path := tvchooser.FileChooser(nil, false, clflashcards_home)
	if path == "" {
		return nil
	}
	err := openFileInEditor(path, *editor, *isNewFile)
	if err != nil {
		return err
	}

	idx.currentCard = path
	return nil

}

func runEditFile(idx *FlashcardIndex, args []string) error {
	editfs := pflag.NewFlagSet("editFile", pflag.ContinueOnError)
	fname := editfs.StringP("filename", "f", "", "Name of file to edit")
	editor := editfs.StringP("editor", "e", "code", "Text editor to open file with")
	isNewFile := editfs.BoolP("newFile", "n", false, "Flag to create new file in same dir as currently selected file if fname does not exist.")

	if err := editfs.Parse(args); err != nil {
		return fmt.Errorf("parsing args: %w", err)
	}

	// Step 1: Resolve the filename (full path)
	filePath, err := resolveFileName(idx, editfs, *fname)
	if err != nil {
		return err
	}
	if filePath == "" {
		fmt.Println(`NO FILE SELECTED. Either use:
    1)  '-f' flag and specify file 
    or 2) specify file without '-f' as first positional arg after 'edit' 
    or 3) use 'select' to change current flashcard.`)
		return nil
	}

	// Step 2: Open and edit the file
	if err := openFileInEditor(filePath, *editor, *isNewFile); err != nil {
		return err
	}

	// Update current flashcard after editing
	idx.currentCard = filePath
	return nil
}

// resolveFileName determines the correct filename to edit
func resolveFileName(idx *FlashcardIndex, editfs *pflag.FlagSet, fname string) (string, error) {
	if !editfs.Changed("filename") {
		// No -f flag: check positional args or fallback to current card
		positionalArgs := editfs.Args()
		if len(positionalArgs) > 0 {
			fname = positionalArgs[0]
		} else {
			fname = idx.currentCard
		}
		if fname == "" {
			return "", nil
		}
	}

	// Filename provided with -f flag
	fname = addTxtExtension(fname)
	fullpath, err := idx.SmartCardFinder(fname)
	if err != nil {
		return "", err
	}
	idx.SetCurrentFile(fullpath)
	return fullpath, nil
}

// openFileInEditor handles file creation/opening and launching the editor
func openFileInEditor(filePath, editor string, isNew bool) error {
	flags := os.O_RDWR
	if isNew {
		flags |= os.O_CREATE
	}

	f, err := os.OpenFile(filePath, flags, 0644)
	if err != nil {
		if os.IsNotExist(err) && !isNew {
			fmt.Printf("Error: File %s does not exist. If you want to create a new file with the edit command use -n flag\n", filePath)
			return nil
		}
		return fmt.Errorf("opening file: %s: %w", filePath, err)
	}
	defer f.Close()

	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running editor: %s: %w", editor, err)
	}

	fmt.Println("Editing finished!")
	return nil
}
