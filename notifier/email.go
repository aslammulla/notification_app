package notifier

import "fmt"

// EmailNotifier is a concrete implementation of the Notifier interface.
// It simulates sending a message via email.
type EmailNotifier struct{}

// Notify sends the message as an email.
func (e EmailNotifier) Notify(message string) error {
	fmt.Println("[Email] Sent message:", message)
	return nil
}
