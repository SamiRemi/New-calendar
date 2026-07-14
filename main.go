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
	e1, err1 := events.NewEvent("Созвон", "2025/07/25 20:35")
	if err1 != nil {
		fmt.Println("Ошибка создания события:", err1)
		return
	}
	e2, err2 := events.NewEvent("Тect", "2026/07/04 ")
	if err2 != nil {
		fmt.Println("Ошибка создания события:", err2)
		return
	}
	calendar.AddEvent("event1", e)
	calendar.AddEvent("event2", e1)
	calendar.AddEvent("event3", e2)
	calendar.DeleteEvent("event2")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")
	fmt.Println("hello")

	err3 := calendar.EditEvent("event11", "Meet", "2025/07/25 20:35")
	if err3 != nil {
		fmt.Println("Ошибка:", err3)
	}
	calendar.ShowEvents()
}
