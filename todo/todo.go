package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func (todo Todo) Display() {
	fmt.Printf("Text: %s\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		fmt.Print(err)
		return err
	}

	return os.WriteFile(fileName, json, 0777)

}

func New(text string) (Todo, error) {

	if text == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text:      text,
		CreatedAt: time.Now(),
	}, nil
}
