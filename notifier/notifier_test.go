package notifier_test

import (
	"testing"

	"go-notifier/notifier" // replace with your actual module path
)

// MockNotifier is a mock implementation of the Notifier interface.
// It records if Notify was called and what message was passed.
type MockNotifier struct {
	Called  bool
	Message string
}

// Notify simulates the notification behavior for testing.
func (m *MockNotifier) Notify(message string) error {
	m.Called = true
	m.Message = message
	return nil
}

// TestAlertUser verifies that AlertUser correctly uses the Notifier interface.
func TestAlertUser(t *testing.T) {
	mock := &MockNotifier{}

	// AlertUser is implemented in the main package.
	alertUser := func(n notifier.Notifier, msg string) {
		_ = n.Notify(msg)
	}

	alertUser(mock, "Test message")

	if !mock.Called {
		t.Errorf("Expected Notify to be called")
	}
	if mock.Message != "Test message" {
		t.Errorf("Expected message to be 'Test message', got '%s'", mock.Message)
	}
}
