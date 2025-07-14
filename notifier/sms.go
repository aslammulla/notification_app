package notifier

import "fmt"

// SMSNotifier is a concrete implementation of the Notifier interface.
// It simulates sending a message via SMS.
type SMSNotifier struct{}

// Notify sends the message as an SMS.
func (s SMSNotifier) Notify(message string) error {
	fmt.Println("[SMS] Sent message:", message)
	return nil
}
