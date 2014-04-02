package main

import (
    config "../config"
    util "../util"
    "flag"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
)

var Db util.SensorDB = util.MakeSensorDB("localhost")
var ReqCount int = 0

func SensorRoute(w http.ResponseWriter, r *http.Request) {
    ReqCount++
    r.ParseForm()

    var sd util.SensorData

    sd.Id = r.Form.Get("id")
    sd.Timestamp = r.Form.Get("timestamp")
    sd.Latitude, _ = strconv.ParseFloat(r.Form.Get("latitude"), 64)
    sd.Longitude, _ = strconv.ParseFloat(r.Form.Get("longitude"), 64)
    sd.Data = r.Form.Get("data")

    Db.Insert(sd)
}

func ClientRoute(w http.ResponseWriter, r *http.Request) {
    ReqCount++
    param_list := strings.Split(r.RequestURI, "/")[2:]

    lat, _ := strconv.ParseFloat(param_list[0], 64)
    lon, _ := strconv.ParseFloat(param_list[1], 64)
    rad, _ := strconv.ParseFloat(param_list[2], 64)

    sd_list := Db.GetNear(lat, lon, rad)
    fmt.Fprint(w, sd_list)
}

func main() {

    UIHandler()
    http.HandleFunc(config.WorkerSensorRoute, util.RouteWrapper(SensorRoute))
    http.HandleFunc(config.WorkerClientRoute, util.RouteWrapper(ClientRoute))

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

    err := http.ListenAndServe(*addr_flag+":"+*port_flag, nil)
    log.Println(err)
}
