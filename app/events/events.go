package events

import (
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

type Event struct {
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("Неверный формат даты")
	}
	return Event{
		Title:   title,
		StartAt: t,
	}, nil
}
func show() {
	fmt.Println("hello")
}
