package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {
	title, content := getNoteDate()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}
	userNote.Display()
}

func getNoteDate() (string, string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
		return ""
	}

	text = strings.TrimPrefix(text, "\n")
	text = strings.TrimPrefix(text, "\r")
	return text
}
