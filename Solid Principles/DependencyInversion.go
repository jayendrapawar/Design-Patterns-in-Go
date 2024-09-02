package main

// class should depend on interfaces rather than concerte class

// The Dependency Inversion Principle states that high-level modules should not depend 
// on low-level modules, but both should depend on abstractions. This principle promotes 
// loose coupling between components, making the code more maintainable and testable.


type EmailService1 struct{}

func (e *EmailService1) Send(to string, message string) {
 // Send email
}

// NotificationService is directly Depended on EmailService
// make it difficult to switch another notification Method
type NotificationService1 struct {
 emailService1 *EmailService1
}

func (n *NotificationService1) Notify(to string, message string) {
 n.emailService1.Send(to, message)
}

//------------------------------------------------------------------------------------------


type MessageSender interface {
    Send(to string, message string)
}

type EmailService struct{}

func (e *EmailService) Send(to string, message string) {
    // Send email
}

type SMSService struct{}

func (s *SMSService) Send(to string, message string) {
    // Send SMS
}

// Now notification service depends on messageSender
type NotificationService struct {
    messageSender MessageSender
}

func (n *NotificationService) Notify(to string, message string) {
    n.messageSender.Send(to, message)
}