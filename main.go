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
	err = userNote.Save()
	if err != nil {
		fmt.Print("Saving data failed!")
		return
	}
	fmt.Println("Saving the note")
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

	text = strings.TrimSpace(text)
	return text
}
