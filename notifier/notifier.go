package notifier

// Notifier is the interface that defines a behavior.
// Any type that implements the Notify method satisfies this interface.
type Notifier interface {
	Notify(message string) error
}
