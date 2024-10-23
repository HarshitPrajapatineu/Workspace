package main

import (
	"bufio"
	"fmt"
	"os"
	"project/note"
	"project/todo"
	"strings"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = OutputData(userNote)

	if err != nil {
		return
	}

	err = OutputData(todo)

	if err != nil {
		return
	}

}

func SaveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Print("Save Failed")
		return err
	}

	fmt.Print("Save Succeeded!")
	return nil
}

func OutputData(data outputtable) error {
	data.Display()
	return SaveData(data)
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")

	content := getUserInput("Note Content: ")

	return title, content

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// var value string
	// fmt.Scanln(&value)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
