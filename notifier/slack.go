package notifier

import "fmt"

// SlackNotifier is a concrete implementation of the Notifier interface.
// It simulates posting a message to a Slack channel.
type SlackNotifier struct{}

// Notify sends the message to Slack.
func (s SlackNotifier) Notify(message string) error {
	fmt.Println("[Slack] Posted message:", message)
	return nil
}
