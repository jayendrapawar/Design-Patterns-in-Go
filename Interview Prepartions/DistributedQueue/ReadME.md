To help visualize the distributed queue system, we'll break it down into tables and flow diagrams to illustrate the key data structures and the message flow within the system.

### Data Structures

1. **Message Struct**
   - Represents a single message.
   - Contains:
     - `Value`: The content of the message.

   | Field | Type  | Description                |
   |-------|-------|----------------------------|
   | Value | string| The content of the message |

2. **Topic Struct**
   - Manages a queue of messages and consumer channels.
   - Contains:
     - `Name`: The name of the topic.
     - `Queue`: A buffered channel of `Message` for storing messages.
     - `Consumers`: A map of consumer IDs to channels for receiving messages.
     - `mu`: A `sync.RWMutex` for managing concurrent access.

   | Field      | Type                            | Description                                    |
   |------------|---------------------------------|------------------------------------------------|
   | Name       | string                          | The name of the topic                          |
   | Queue      | chan Message                    | Channel to hold messages                       |
   | Consumers  | map[string]chan Message         | Map of consumer IDs to channels                |
   | mu         | sync.RWMutex                     | Mutex for concurrent access                   |

3. **QueueSystem Struct**
   - Manages multiple topics.
   - Contains:
     - `Topics`: A map of topic names to `Topic` instances.
     - `mu`: A `sync.RWMutex` for managing concurrent access.

   | Field   | Type                             | Description                            |
   |---------|----------------------------------|----------------------------------------|
   | Topics  | map[string]*Topic                | Map of topic names to Topic instances  |
   | mu      | sync.RWMutex                     | Mutex for concurrent access            |

4. **Producer Struct**
   - Represents a producer that publishes messages.
   - Contains:
     - `ID`: The ID of the producer.

   | Field | Type  | Description             |
   |-------|-------|-------------------------|
   | ID    | string| The ID of the producer  |

5. **Consumer Struct**
   - Represents a consumer that subscribes to topics.
   - Contains:
     - `ID`: The ID of the consumer.

   | Field | Type  | Description             |
   |-------|-------|-------------------------|
   | ID    | string| The ID of the consumer  |

### Flow Diagram

Here's a visual flow of how messages move through the system:

#### 1. Topic Creation

- **`QueueSystem`** → **`NewTopic`** → **`Topic`**

#### 2. Message Publishing

- **`Producer`** → **`Publish`** → **`Topic.Queue`**

#### 3. Message Distribution

- **`Topic.Queue`** → **`Distribute`** → **`Topic.Consumers`**

#### 4. Message Consumption

- **`Topic.Consumers`** → **`Consumer`** → **`Process Message`**

### Flow Diagram Description

1. **Topic Creation:**
   - A `QueueSystem` creates a new `Topic` using `NewTopic`. The `Topic` is added to the `QueueSystem`'s `Topics` map.

2. **Message Publishing:**
   - A `Producer` publishes a `Message` to a `Topic` using the `Publish` method. The message is sent to the `Topic.Queue` channel.

3. **Message Distribution:**
   - The `Topic` continuously reads messages from `Topic.Queue` and distributes them to all channels in `Topic.Consumers` using the `Distribute` method.

4. **Message Consumption:**
   - Each `Consumer` receives messages from its associated channel in `Topic.Consumers` and processes them asynchronously.

### Summary

- **Producer** → **Topic.Queue**: Producers send messages to the topic's queue.
- **Topic.Queue** → **Topic.Consumers**: The topic distributes messages to consumer channels.
- **Consumer** → **Process Message**: Consumers receive and process messages from their channels.

These visualizations and flow diagrams should help clarify how the distributed queue system operates and how data flows through the various components.