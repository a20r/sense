package main

import (
    config "../config"
    "net/http"
    "strconv"
    "time"
)

func UpdateWorkers(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    address := r.Form.Get("address")
    req_count, _ := strconv.Atoi(r.Form.Get("count"))
    current_time := time.Now().Unix()

    load_data, in_map := heartbeatMap[address]
    if !in_map {
        heartbeatMap[address] = LoadData{current_time, req_count}
    } else {
        load_data.Timestamp = current_time
        load_data.Frequency = req_count
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
    }
}
