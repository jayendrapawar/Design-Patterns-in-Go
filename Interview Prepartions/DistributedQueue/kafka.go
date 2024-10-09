package main

import (
	"fmt"
	"sync"
	"time"
)

/*  CODE APPROACH */
// Message Struct: Represents individual messages.
// Topic Struct: Manages a queue of messages and consumer channels.
// QueueSystem Struct: Oversees multiple topics and manages topic creation and retrieval.
// AddConsumer: Adds a consumer to a topic with its own channel.
// Publish: Sends messages to a topic’s queue.
// Distribute: Sends messages from the queue to all consumers.
// Producer: Publishes messages to topics.
// Consumer: Subscribes to topics and processes messages asynchronously.
// Concurrency Handling: Uses sync.RWMutex to manage concurrent access to topics and consumers.
// Backpressure Handling: Logs missed messages if consumer channels are full.

// Message represents a message in the queue
type Message struct {
	Value string
}

// Topic represents a single topic with a queue of messages
type Topic struct {
	Name      string
	Queue     chan Message
	Consumers map[string]chan Message
	mu        sync.RWMutex
}

// NewTopic creates a new topic
func NewTopic(name string) *Topic {
	return &Topic{
		Name:      name,
		Queue:     make(chan Message, 100),
		Consumers: make(map[string]chan Message),
	}
}

// AddConsumer adds a new consumer to the topic
func (t *Topic) AddConsumer(consumerID string) chan Message {
	t.mu.Lock()
	defer t.mu.Unlock()
	consumerChannel := make(chan Message, 100)
	t.Consumers[consumerID] = consumerChannel
	return consumerChannel
}

// Publish publishes a message to the topic
func (t *Topic) Publish(message Message) {
	t.Queue <- message
}

// Distribute distributes messages to all consumers
func (t *Topic) Distribute() {
	for msg := range t.Queue {
		t.mu.RLock()
		for consumerID, consumerChan := range t.Consumers {
			select {
			case consumerChan <- msg:
				fmt.Printf("%s received %s\n", consumerID, msg.Value)
			default:
				fmt.Printf("%s missed message %s due to a full buffer\n", consumerID, msg.Value)
			}
		}
		t.mu.RUnlock()
	}
}

// QueueSystem manages multiple topics
type QueueSystem struct {
	Topics map[string]*Topic
	mu     sync.RWMutex
}

// NewQueueSystem creates a new queue system
func NewQueueSystem() *QueueSystem {
	return &QueueSystem{
		Topics: make(map[string]*Topic),
	}
}

// CreateTopic creates a new topic in the queue system
func (qs *QueueSystem) CreateTopic(name string) {
	qs.mu.Lock()
	defer qs.mu.Unlock()
	if _, exists := qs.Topics[name]; !exists {
		topic := NewTopic(name)
		qs.Topics[name] = topic
		go topic.Distribute()
	}
}

// GetTopic retrieves a topic by name
func (qs *QueueSystem) GetTopic(name string) *Topic {
	qs.mu.RLock()
	defer qs.mu.RUnlock()
	return qs.Topics[name]
}

// Producer publishes messages to topics
type Producer struct {
	ID string
}

// Publish publishes a message to a topic
func (p *Producer) Publish(topic *Topic, message string) {
	topic.Publish(Message{Value: message})
}

// Consumer subscribes to topics and consumes messages
type Consumer struct {
	ID string
}

// Subscribe subscribes the consumer to a topic
func (c *Consumer) Subscribe(topic *Topic) {
	consumerChan := topic.AddConsumer(c.ID)
	go func() {
		for msg := range consumerChan {
			fmt.Printf("%s received %s\n", c.ID, msg.Value)
		}
	}()
}

func main() {
	queueSystem := NewQueueSystem()

	// Create topics
	queueSystem.CreateTopic("topic1")
	queueSystem.CreateTopic("topic2")

	// Create producers
	producer1 := &Producer{ID: "producer1"}
	producer2 := &Producer{ID: "producer2"}

	// Create consumers
	consumer1 := &Consumer{ID: "consumer1"}
	consumer2 := &Consumer{ID: "consumer2"}

	// Get topics
	topic1 := queueSystem.GetTopic("topic1")
	topic2 := queueSystem.GetTopic("topic2")

	// Consumers subscribe to topics
	consumer1.Subscribe(topic1)
	consumer1.Subscribe(topic2)
	consumer2.Subscribe(topic1)

	// Producers publish messages to topics
	producer1.Publish(topic1, "message 1 to topic1")
	producer1.Publish(topic2, "message 1 to topic2")
	producer2.Publish(topic1, "message 2 to topic1")

	// Give some time for messages to be processed
	time.Sleep(1 * time.Second)
}
