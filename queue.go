/*
    Simple fixed-size FIFO message queue. The main()
    function tests out the PushMessage() and PopMessage()
    functionality through goroutines. 

    The queue implements an integer  channel that acts like
    a semaphore. PopMessage() waits till there is something
    in the queue to pop. PushMessage() writes a value to the
    semaphore when a new message is added to the queue.
*/
package main

import (
    "fmt"
    "time"
)

type Message struct {
    contents string
}

type MessageQueue struct {
    messages []Message
    max_size, used int
    sem chan int
}

func CreateMQueue (q_size int) MessageQueue {
    m:= MessageQueue{ max_size: q_size,
                      used: 0,
                      sem: make(chan int, q_size) }
    return m
}

func (q *MessageQueue) PopMessage() interface{} {
    if (q.used == 0) {
        return nil
    }

    // Wait till there is something to pop
    <-q.sem

    // Pop and update used value
    index := q.used - 1
    m := q.messages[index]
    q.messages = q.messages[0:index]
    q.used -= 1

    fmt.Println("Removing", m.contents)
    return m
}

func (q *MessageQueue) PushMessage(msg string) bool {
    // Infinite waiting loop if queue is full
    for (q.used == q.max_size) {}

    // Add message to queue and update used value
    m := Message {contents: msg}
    fmt.Println("Adding", msg)
    q.messages = append(q.messages, m)
    q.used += 1

    // Signal that there is a new message
    q.sem <- 1
    return true
}

func main() {
    q := CreateMQueue(2)
    q.PushMessage("Msg 1")
    q.PushMessage("Hello")
    go func() {
        q.PopMessage()
        time.Sleep(200*time.Millisecond)
        q.PopMessage()
        q.PushMessage("News")
        fmt.Println("First", q.messages)
    }()
    go func() {
        q.PushMessage("Music")
        q.PopMessage()
        q.PushMessage("Sport")
        fmt.Println("Second", q.messages)
    }()
    fmt.Println(q.messages, q.used, len(q.sem))
    time.Sleep(5000*time.Millisecond)
}
