package events

import (
	"errors"
	"fmt"
	"time"

	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
)

type Event struct {
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	if !validation.IsValidTitle(title) {
		return Event{}, fmt.Errorf("неверный формат заголовка %q", title)
	}
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return Event{}, errors.New("Неверный формат даты")
	}
	return Event{
		Title:   title,
		StartAt: t,
	}, nil
}

func (e *Event) Update(title string, dateStr string) error {
	isValid := validation.IsValidTitle(title)
	if !isValid {
		return fmt.Errorf("неверный формат заголовка %q", title)
	}
	time, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return fmt.Errorf("Неверный формат даты %q", time)
	}

	e.Title = title
	e.StartAt = time
	return nil
}
