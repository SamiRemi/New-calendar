package main

import (
	"fmt"

	"github.com/SamiRemi/project/app/calendar"

	"github.com/SamiRemi/project/app/events"
)

func main() {
	e, err := events.NewEvent("Встреча", "2025/06/12 16:33")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}
	calendar.AddEvent("event1", e)
	calendar.ShowEvents()
}
