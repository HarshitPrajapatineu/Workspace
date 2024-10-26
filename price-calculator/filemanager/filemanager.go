package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	Path           string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.Path)

	if err != nil {
		fmt.Println("Error Opening File")
		fmt.Println(err)
		return nil, errors.New("Error Opening File")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading file content Failed")
		fmt.Println(err)
		return nil, errors.New("Reading file content Failed")
	}

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Error Creating File.")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(data)

	time.Sleep(3 * time.Second)

	if err != nil {
		return errors.New("Error Converting to JSON.")
	}

	return nil
}

func New(path, outputFilePath string) FileManager {
	return FileManager{
		Path:           path,
		OutputFilePath: outputFilePath,
	}
}
