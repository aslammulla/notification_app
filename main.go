package main

import (
	"fmt"
	"go-notifier/notifier" // replace with your actual module path
)

// AlertUser is a reusable function that accepts any Notifier interface.
// This allows us to send alerts via email, SMS, Slack, etc., without changing this function.
func AlertUser(n notifier.Notifier, message string) {
	err := n.Notify(message)
	if err != nil {
		fmt.Println("Notification failed:", err)
	}
}

func main() {
	// Create instances of each notifier
	email := notifier.EmailNotifier{}
	sms := notifier.SMSNotifier{}
	slack := notifier.SlackNotifier{}

	// Send notifications via different channels
	AlertUser(email, "Welcome via Email!")
	AlertUser(sms, "Your OTP is 123456")
	AlertUser(slack, "Deployment completed successfully!")
}
