# clflashcards

Simple flashcard tool entirely based in the terminal. Mouse free experience. 

[![Demo Video](https://img.shields.io/badge/▶️%20Watch%20Demo-blue)](https://drive.google.com/file/d/1EujvxUU096P6VJBj1O9cBhrfvKwCMiqD/view?usp=drive_link)


## Description

Flashcard files are simply .txt files where each line contains a question and an answer seperated by | e.g. What does GPT stand for? | Generative Pre-Trained Transformer
...

## Getting Started

### Dependencies

* ...


### Installing

```
go install github.com/samirm-git/clflashcards@latest
```
Or download zip file from releases tab

**Need to add the exe file to PATH after installing**

### Commands

```
create [--question] [--answer]

createFile <path>

edit [--filename] [--editor] [--isNewFile]

select <path>

selectgui

show [n]

testme [--random] [--numberOfQuestions]

help [command]

refresh
```


## Version History

* 1.1.1
  * Cleanup
* 1.1.0
   * Improved REPL with arrow key scrolling
* 1.0.0
    * Initial Release


## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE License - see the LICENSE.md file for details

## Acknowledgments

* [tvchooser](https://github.com/AEROGU/tvchooser/)
* [readline](https://github.com/chzyer/readline)
