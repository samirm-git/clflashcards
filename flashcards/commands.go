package flashcards

type Command struct {
	Name        string
	Description string
	Run         func(idx *FlashcardIndex, args []string) error
	RunHelp     func() // optional helper function for showing command-specific help
}

var commandsMap = map[string]*Command{
	"create": {
		Name:        "create",
		Description: "Create a new flashcard",
		Run:         runCreate,
		RunHelp:     helpCreate,
	},
	"show": {
		Name:        "show",
		Description: "Show a flashcard",
		Run:         runShow,
		RunHelp:     helpShow,
	},
	"edit": {
		Name:        "edit",
		Description: "Edit a flashcard file",
		Run:         runEditFile,
		RunHelp:     helpEdit,
	},
	"editgui": {
		Name:        "editgui",
		Description: "Edit a flashcard in GUI mode",
		Run:         runEditFileGUI,
		RunHelp:     helpEditGui,
	},
	"testme": {
		Name:        "testme",
		Description: "Test yourself with flashcards",
		Run:         runTestme,
		RunHelp:     helpTestMe,
	},
	"refresh": {
		Name:        "refresh",
		Description: "build flashcard file index again",
		Run:         runRefresh,
		RunHelp:     helpReresh,
	},
	"select": {
		Name:        "select",
		Description: "Select a flashcard (or open GUI if no args)",
		Run: func(idx *FlashcardIndex, args []string) error {
			if len(args) == 0 {
				return runSelectFlashCardGUI(idx)
			}
			return runSelectFlashCard(idx, args[0])
		},
		RunHelp: helpSelect,
	},
}

// Add "help" command after all others
func init() {
	commandsMap["help"] = &Command{
		Name:        "help",
		Description: "Show available commands and usage information",
		Run:         runHelp,
		RunHelp:     helpHelp,
	}
}

func GetCommand(name string) (*Command, bool) {
	cmd, ok := commandsMap[name]
	return cmd, ok
}

func AllCommands() map[string]*Command {
	return commandsMap
}
