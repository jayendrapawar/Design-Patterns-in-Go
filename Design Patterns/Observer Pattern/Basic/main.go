package main

import "fmt"

// Publisher interface defines the methods that a Publisher should have
type Publisher interface {
	Add(sub Subscriber)
	Remove(sub Subscriber)
	NotifyAll()
	SetData(msg string)
	GetData() string
}

// Subscriber interface defines the method that a Subscriber should have
type Subscriber interface {
	Update(msg string)
}

// Concrete Publisher implementation
type publisher struct {
	subscriberList []Subscriber
	data           string
}

// Constructor for a new publisher
func GetNewPublisher() Publisher {
	return &publisher{
		subscriberList: make([]Subscriber, 0),
	}
}

// Adds a Subscriber to the Publisher's list
func (pub *publisher) Add(sub Subscriber) {
	pub.subscriberList = append(pub.subscriberList, sub)
}

// Removes a Subscriber from the Publisher's list
func (pub *publisher) Remove(sub Subscriber) {
	for i, s := range pub.subscriberList {
		if s == sub {
			pub.subscriberList = append(pub.subscriberList[:i], pub.subscriberList[i+1:]...)
			break
		}
	}
}

// Notifies all Subscribers of an update
func (pub *publisher) NotifyAll() {
	for _, sub := range pub.subscriberList {
		sub.Update(pub.data)
	}
}

// Sets the data in the Publisher
func (pub *publisher) SetData(msg string) {
	pub.data = msg
	pub.NotifyAll()
}

// Gets the current data from the Publisher
func (pub *publisher) GetData() string {
	return pub.data
}

// Concrete Subscriber implementation
type subscriber struct {
	subscriberID string
}

// Constructor for a new subscriber
func GetNewSubscriber(ID string) Subscriber {
	return &subscriber{
		subscriberID: ID,
	}
}

// Updates the Subscriber with new data
func (s *subscriber) Update(msg string) {
	fmt.Printf("Subscriber %s received message: %s\n", s.subscriberID, msg)
}

func main() {
	// Create a Publisher
	pub := GetNewPublisher()

	// Create Subscribers
	sub1 := GetNewSubscriber("1")
	sub2 := GetNewSubscriber("2")

	// Add Subscribers to the Publisher
	pub.Add(sub1)
	pub.Add(sub2)

	// Set data and notify all Subscribers
	pub.SetData("Hello, Subscribers!")

	// Remove a Subscriber and set new data
	pub.Remove(sub1)
	pub.SetData("Another message for remaining Subscribers")
}
