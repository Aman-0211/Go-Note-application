package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	contnet   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("for your reference %v note have this content: \n\n%v", note.title, note.contnet)
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input")
	}

	return Note{
		title:   title,
		contnet: content,
	}, nil
}
