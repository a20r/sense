package main

import (
    util "../util"
    "net/http"
)

func UIHandler() {
    staticHandler := util.FileResponseCreator("static")
    http.HandleFunc("/", util.FileResponseCreator("templates"))
    http.HandleFunc("/css/", staticHandler)
    http.HandleFunc("/js/", staticHandler)
    http.HandleFunc("/img/", staticHandler)
    http.HandleFunc("/favicon.ico", util.FileResponseCreator("static/img"))
}
