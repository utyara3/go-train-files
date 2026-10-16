package main

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Notify(message string) error
}

type BaseNotifier struct {
	MaxNotifies   int
	NotifyHistory []string
}

type ConsoleNotifier struct {
	ConsoleName string
	BaseNotifier
}

type EmailNotifier struct {
	UserEmail string
	BaseNotifier
}

var ErrNotifiesLimitationError = errors.New("reached max amount of notifies")

func checkMaxNotifies(notifier *BaseNotifier) error {
	if len(notifier.NotifyHistory) >= notifier.MaxNotifies {
		return fmt.Errorf("notify error: %w\n", ErrNotifiesLimitationError)
	}
	return nil
}

func (c *ConsoleNotifier) Notify(message string) error {
	if err := checkMaxNotifies(&c.BaseNotifier); err != nil {
		return err
	}

	fmt.Printf("Console Notify: %s\n", message)
	c.NotifyHistory = append(c.NotifyHistory, message)
	return nil
}

func (e *EmailNotifier) Notify(message string) error {
	if err := checkMaxNotifies(&e.BaseNotifier); err != nil {
		return err
	}

	fmt.Printf("Email Notify: %s\n", message)
	e.NotifyHistory = append(e.NotifyHistory, message)
	return nil
}

func main() {
	not1 := &ConsoleNotifier{
		ConsoleName: "BaseConsole",
		BaseNotifier: BaseNotifier{
			MaxNotifies: 10,
		},
	}

	not2 := &EmailNotifier{
		UserEmail: "test@example.com",
		BaseNotifier: BaseNotifier{
			MaxNotifies: 100,
		},
	}

	notifs := []Notifier{not1, not2}

	for i := range notifs {
		notifs[i].Notify("Hello!")
	}
}
