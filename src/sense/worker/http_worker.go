package main

import (
    config "../config"
    "flag"
    "fmt"
    "net/http"
)

func SensorRoute(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Worker")
}

func main() {

    UIHandler()
    http.HandleFunc(config.WorkerSensorRoute, SensorRoute)

    var addr_flag = flag.String(
        "addr",
        "localhost",
        "Address the http server binds to",
    )

    var port_flag = flag.String(
        "port",
        "8000",
        "Port used for http server",
    )

    flag.Parse()

    heart := CreateHeart(1000, "http://"+*addr_flag+":"+*port_flag)
    defer heart.Stop()
    heart.Start()
    http.ListenAndServe(*addr_flag+":"+*port_flag, nil)
}
