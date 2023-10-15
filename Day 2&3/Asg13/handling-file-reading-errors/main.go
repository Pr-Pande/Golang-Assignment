package main

import (
	"fmt"
	"os"
	"strings"
)

func CountWords(content string) (int, error) {
	
	content = strings.TrimSpace(content)
	if content == "" {
		return 0, fmt.Errorf("content is empty or contains only whitespace")
	}

	words := strings.Fields(content)
	return len(words), nil
}

func main() {
	fileContent, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	contentStr := string(fileContent)

	wordCount, err := CountWords(contentStr)
	if err != nil {
		fmt.Printf("Error counting words: %v\n", err)
		return
	}

	fmt.Printf("Number of words in the file: %d\n", wordCount)

	err = os.Remove("sample.txt")
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
	}
}