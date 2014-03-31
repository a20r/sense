package main

import (
    config "../config"
    "log"
    "net/http"
    "time"
)

func UpdateWorkers(w http.ResponseWriter, r *http.Request) {
    log.Println(":: POST --> /heartbeat")
    r.ParseForm()

    address := r.Form.Get("address")
    current_time := time.Now().Unix()

    load_data, in_map := heartbeatMap[address]
    if !in_map {
        heartbeatMap[address] = LoadData{current_time, 0}
    } else {
        load_data.Timestamp = current_time
        heartbeatMap[address] = load_data
    }
}

func removeDeadWorkers(workerMap map[string]LoadData) {
    current_time := time.Now().Unix()
    for addr, load_data := range workerMap {
        if current_time-load_data.Timestamp > config.WorkerTimeRemove {
            delete(workerMap, addr)
        }
    }
}

func removeDeadWorkersLoop(workerMap map[string]LoadData) {
    for {
        time.Sleep(config.TimeDelayRemoveCheck)
        removeDeadWorkers(workerMap)
        log.Println(workerMap)
    }
}
