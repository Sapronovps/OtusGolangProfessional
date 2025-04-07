package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	result := make(Environment)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(dir, entry.Name())
		value, needRemove, err := processFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error processing file %q: %w", filePath, err)
		}

		result[entry.Name()] = EnvValue{
			Value:      value,
			NeedRemove: needRemove,
		}
	}

	return result, nil
}

func processFile(path string) (string, bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", false, fmt.Errorf("error open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	firstLine := scanner.Text()
	if firstLine == "" {
		return "", true, nil
	}

	firstLine = strings.ReplaceAll(firstLine, "\x00", "\n")
	firstLine = strings.TrimRight(firstLine, "\n ")

	return firstLine, false, nil
}
