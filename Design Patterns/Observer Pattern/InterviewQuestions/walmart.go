package main

import "fmt"

// StockPublisher interface defines the methods for the publisher
type StockPublisher interface {
	Add(sub Subscriber)
	Remove(sub Subscriber)
	NotifyAll()
	SetData(msg string)
	GetData() string
}

// Subscriber interface defines the method for subscribers
type Subscriber interface {
	Update(msg string)
}

// Concrete implementation of StockPublisher
type publisher struct {
	subscriberList []Subscriber
	data           string
}

// Constructor for a new publisher
func GetStockPublisher() *publisher {
	return &publisher{
		subscriberList: make([]Subscriber, 0),
	}
}

// Adds a Subscriber to the list
func (pub *publisher) Add(sub Subscriber) {
	pub.subscriberList = append(pub.subscriberList, sub)
}

// Removes a Subscriber from the list
func (pub *publisher) Remove(sub Subscriber) {
	pub.subscriberList = deleteElement(pub.subscriberList, sub)
}

// Helper function to remove an element from a slice
func deleteElement(slice []Subscriber, sub Subscriber) []Subscriber {
	for i, val := range slice {
		if val == sub {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

// Notifies all Subscribers
func (pub *publisher) NotifyAll() {
	for _, sub := range pub.subscriberList {
		sub.Update(pub.data)
	}
}

// Sets the data and notifies all Subscribers
func (pub *publisher) SetData(msg string) {
	pub.data = msg
	pub.NotifyAll()
}

// Gets the current data
func (pub *publisher) GetData() string {
	return pub.data
}

// Concrete implementation of Subscriber for email
type EmailSubscriber struct {
	eSubID string
}

// Constructor for a new EmailSubscriber
func GetNewEmailSubscriber(id string) *EmailSubscriber {
	return &EmailSubscriber{
		eSubID: id,
	}
}

// Update method for EmailSubscriber
func (e *EmailSubscriber) Update(msg string) {
	fmt.Printf("EmailSubscriber %s received message: %s\n", e.eSubID, msg)
}

// Concrete implementation of Subscriber for SMS
type SMSSubscriber struct {
	sSubID string
}

// Constructor for a new SMSSubscriber
func GetNewSMSSubscriber(id string) *SMSSubscriber {
	return &SMSSubscriber{
		sSubID: id,
	}
}

// Update method for SMSSubscriber
func (s *SMSSubscriber) Update(msg string) {
	fmt.Printf("SMSSubscriber %s received message: %s\n", s.sSubID, msg)
}

func main() {
	// Create a StockPublisher
	st := GetStockPublisher()

	// Create Subscribers
	s1 := GetNewEmailSubscriber("1")
	s2 := GetNewSMSSubscriber("2")
	s3 := GetNewEmailSubscriber("3")

	// Add Subscribers to the Publisher
	st.Add(s1)
	st.Add(s2)
	st.Add(s3)

	// Set data and notify all Subscribers
	st.SetData("Hello")

	// Remove a Subscriber and notify remaining Subscribers
	st.Remove(s3)
	st.NotifyAll()

	// Set new data and notify all Subscribers
	st.SetData("Jay")
}
