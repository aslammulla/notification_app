# 📣 Go Notifier Interface Example

A simple and extensible notification system built using Go interfaces.  
Supports multiple types of notifications like **Email**, **SMS**, and **Slack** — all through a common interface.

> This project demonstrates how to use **Go interfaces** to build flexible, testable, and maintainable software.

---

## 📦 Project Structure

```go
notifier/
├── notifier.go # Interface definition
├── email.go # EmailNotifier implementation
├── sms.go # SMSNotifier implementation
├── slack.go # SlackNotifier implementation
├── notifier_test.go # Unit test using a mock implementation
main.go # Entry point demonstrating usage

```

---

## 📘 What’s Inside?

### Notifier Interface

Defines the behavior for all notifiers:

```go
type Notifier interface {
    Notify(message string) error
}
```

### Concrete Implementations

Each notifier (Email, SMS, Slack) implements the Notifier interface independently:

```go
func (e EmailNotifier) Notify(message string) error
func (s SMSNotifier) Notify(message string) error
func (s SlackNotifier) Notify(message string) error
```

### How to Run

#### Clone the repo
```bash
git clone https://github.com/your-username/go-notifier.git
cd go-notifier
```
#### Initialize Go modules
```bash
go mod init go-notifier
go mod tidy
```

#### Run the app
```bash
go run main.go
```
#### Example Output
```bash
[Email] Sent message: Welcome via Email!
[SMS] Sent message: Your OTP is 123456
[Slack] Posted message: Deployment completed successfully!
```