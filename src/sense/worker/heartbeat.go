package main

import (
    // "fmt"
    config "../config"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

type Heart struct {
    Interval    time.Duration
    LocalUrl    string
    StopChannel chan int
}

//creates a heart that will be used to send "beats" to the server
func CreateHeart(interval time.Duration, local_url string) *Heart {
    stop_channel := make(chan int)
    return &Heart{interval, local_url, stop_channel}
}

//send an error
func (hb *Heart) Beat() error {
    _, err := http.PostForm(
        config.BrokerHeartbeatUrl,
        url.Values{"address": {hb.LocalUrl}, "count": {strconv.Itoa(ReqCount)}},
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
