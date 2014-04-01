package main

import (
    // config "../config"
    util "../util"
    "flag"
    "fmt"
    "net/http"
)

type LoadData struct {
    Timestamp int64
    Frequency int
}

var heartbeatMap map[string]LoadData = make(map[string]LoadData)

func getMinLoad() string {
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

    return *minAddr
}

func GetURL(w http.ResponseWriter, r *http.Request) {
    if len(heartbeatMap) < 1 {
        panic("No workers")
    }

    minAddr := getMinLoad()
    load_data := heartbeatMap[minAddr]
    load_data.Frequency++
    heartbeatMap[minAddr] = load_data

    fmt.Fprint(w, util.Response{"address": minAddr})
}

func MobileDeviceReroute(w http.ResponseWriter, r *http.Request) {
    if len(heartbeatMap) < 1 {
        panic("No workers")
    }

    minAddr := getMinLoad()

    load_data := heartbeatMap[minAddr]
    load_data.Frequency++
    heartbeatMap[minAddr] = load_data

    http.Redirect(w, r, minAddr+"/temp.html", http.StatusFound)
}

func main() {

    http.HandleFunc("/heartbeat", util.RouteWrapper(UpdateWorkers))
    http.HandleFunc("/register", util.RouteWrapper(MobileDeviceReroute))
    http.HandleFunc("/url", util.RouteWrapper(GetURL))

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
