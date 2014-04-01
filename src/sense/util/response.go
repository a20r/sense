package util

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
)

type Page struct {
    Title string
    Body  []byte
}

type Response map[string]interface{}

func (r Response) String() (s string) {
    b, err := json.Marshal(r)
    if err != nil {
        s = ""
        return
    }
    s = string(b)
    return
}

func loadPage(folder, title string) (*Page, error) {
    filename := folder + "/" + title
    body, err := ioutil.ReadFile(filename)

    if err != nil {
        return nil, err
    }

    return &Page{Title: title, Body: body}, nil
}

func FileResponseCreator(folder string) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        var p *Page

        if len(r.URL.Path) == 1 {
            p, _ = loadPage("templates", "index.html")
        } else {
            p, _ = loadPage(folder, r.URL.Path[1:])
        }

        if p != nil {
            w.Write(p.Body)
        }
    }
}

func RouteWrapper(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
    ret_func := func(w http.ResponseWriter, r *http.Request) {
        log.Println(":: " + r.Method + " --> " + r.RequestURI)
        handler(w, r)
    }

    return ret_func
}
