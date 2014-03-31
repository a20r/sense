package main

import (
    "flag"
    "net/http"
)

func Reroute(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "http://localhost:8001/job", http.StatusFound)
}

func main() {

    http.HandleFunc("/helloworld", Reroute)

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

    http.ListenAndServe(*addr_flag+":"+*port_flag, nil)
}
