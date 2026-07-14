package calendar

import (
	"fmt"

	"github.com/SamiRemi/project/app/events"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) {
	eventsMap[key] = e
	fmt.Println("Событие добавлено:", e.Title)
}

func ShowEvents() {
	for _, e := range eventsMap {
		fmt.Printf("\nСобытие: %s || Дата:%s\n ", e.Title, e.StartAt.Format("2006-01-02 15:04"))
	}
}
