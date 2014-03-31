package main

import (
    // "fmt"
    config "../config"
    "net/http"
    "net/url"
    "time"
)

type Heart struct {
    Interval    time.Duration
    LocalUrl    string
    StopChannel chan int
}

func CreateHeart(interval time.Duration, local_url string) *Heart {
    stop_channel := make(chan int)
    return &Heart{interval, local_url, stop_channel}
}

func (hb *Heart) Beat() error {
    _, err := http.PostForm(
        config.BrokerHeartbeatUrl,
        url.Values{"address": {hb.LocalUrl}},
    )
    return err
}

func (hb *Heart) Start() {
    time_func := func() {
        for {
            select {
            case <-hb.StopChannel:
                return
            default:
                hb.Beat()
                time.Sleep(hb.Interval * time.Millisecond)
            }
        }
    }

    go time_func()
}

func (hb *Heart) Stop() {
    hb.StopChannel <- 1
}
