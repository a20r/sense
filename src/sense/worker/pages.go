package main

import (
    util "../util"
    "net/http"
)

func UIHandler() {
    staticHandler := util.FileResponseCreator("static")
    http.HandleFunc("/", util.FileResponseCreator("static/html"))
    http.HandleFunc("/css/", staticHandler)
    http.HandleFunc("/js/", staticHandler)
    http.HandleFunc("/images/", staticHandler)
    http.HandleFunc("/favicon.ico", util.FileResponseCreator("static/img"))
}
