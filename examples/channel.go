package main

import (
    "fmt"
    "github.com/timonv/pusher"
    "time"
)

func main() {
    client := pusher.NewClient("appid", "key", "secret", false)

    done := make(chan bool)

    go func() {
        channel, err := client.Channel("common", nil)
        if err != nil {
            fmt.Printf("Error %s\n", err)
        } else {
            fmt.Println(channel)
        }
        done <- true
    }()

    select {
    case <-done:
        fmt.Println("Done :-)")
    case <-time.After(1 * time.Minute):
        fmt.Println("Timeout :-(")
    }
}
