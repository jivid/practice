package main

import (
    "fmt"
    "queue"
    "logging"
    "math/rand"
)

func main() {
    logfile := logging.CreateLogger("queue.log", 3)

    q := queue.CreateMQueue(3)
    for i := 0; i < 4; i++ {
        r := rand.Int()
        m := fmt.Sprintf("Pushing %d onto queue", r)
        logfile.Output(2, m)
        q.PushMessage(string(r))
    }
}
