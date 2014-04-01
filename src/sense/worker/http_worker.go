package main

import (
    config "../config"
    util "../util"
    "flag"
    "net/http"
    "strconv"
)

var Db util.SensorDB = util.MakeSensorDB("localhost")

func SensorRoute(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    var sd util.SensorData

    sd.Id = r.Form.Get("id")
    sd.Timestamp = r.Form.Get("timestamp")
    sd.Latitude, _ = strconv.ParseFloat(r.Form.Get("latitude"), 64)
    sd.Longitude, _ = strconv.ParseFloat(r.Form.Get("longitude"), 64)
    sd.Data = r.Form.Get("data")

    Db.Insert(sd)
}

func main() {

    UIHandler()
    http.HandleFunc(config.WorkerSensorRoute, util.RouteWrapper(SensorRoute))

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
