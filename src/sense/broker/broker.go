package main

import (
    // config "../config"
    "flag"
    "net/http"
)

type LoadData struct {
    Timestamp int64
    Frequency int
}

var heartbeatMap map[string]LoadData = make(map[string]LoadData)

func MobileDeviceReroute(w http.ResponseWriter, r *http.Request) {
    if len(heartbeatMap) < 1 {
        panic("No workers")
    }

    var minLoadData *LoadData = nil
    var minAddr *string = nil

    for addr_key, load_data := range heartbeatMap {
        if minLoadData == nil {
            minLoadData = new(LoadData)
            *minLoadData = load_data

            minAddr = new(string)
            *minAddr = addr_key
        }

        if load_data.Frequency < minLoadData.Frequency {
            *minLoadData = load_data
            *minAddr = addr_key
        }
    }
    var addr string
    if minAddr == nil {
        for addr_key, _ := range heartbeatMap {
            addr = addr_key
            break
        }
    } else {
        addr = *minAddr
        minLoadData.Frequency++
    }

    load_data := heartbeatMap[addr]
    load_data.Frequency++
    heartbeatMap[addr] = load_data

    http.Redirect(w, r, addr+"/temp.html", http.StatusFound)
}

func main() {

    http.HandleFunc("/heartbeat", UpdateWorkers)
    http.HandleFunc("/register", MobileDeviceReroute)

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

    go removeDeadWorkersLoop(heartbeatMap)

    http.ListenAndServe(*addr_flag+":"+*port_flag, nil)
}
