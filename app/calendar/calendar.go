package calendar

import (
	"fmt"

	"github.com/SamiRemi/project/app/events"
	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
)

var calendarEvents = make(map[string]events.Event)

func AddEvent(title string, date string) (events.Event, error) {
	e, err := events.NewEvent(title, date)
	if err != nil {
		return e, fmt.Errorf("Ошибка создание события %s\n", title)
	}

	calendarEvents[e.ID] = e
	return e, nil
}

func ShowEvents() {
	for _, e := range calendarEvents {
		fmt.Printf("\nСобытие: %s || Дата:%s\n ", e.Title, e.StartAt.Format("2006-01-02 15:04"))
	}
}

func DeleteEvent(title string) error {
	_, exists := calendarEvents[title]
	if !exists {
		return fmt.Errorf("Событие с заголовком %q не найдено\n", title)
	}
	delete(calendarEvents, title)
	fmt.Printf("Событие %q успешно удалено\n", title)
	return nil
}

func EditEvent(key, title, startAt string) error {
	e, exists := calendarEvents[key]
	if !exists {
		return fmt.Errorf("событие с ключом %q не найдено", key)
	}

	if !validation.IsValidTitle(title) {
		return fmt.Errorf("неверный формат заголовка %q", title)
	}

	t, err := dateparse.ParseAny(startAt)
	if err != nil {
		return fmt.Errorf("не удалось распарсить дату %q: %w", startAt, err)
	}

	e.Title = title
	e.StartAt = t

	calendarEvents[key] = e

	return nil
}
