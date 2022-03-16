package main

import (
	"fmt"
	"net/http"
)

func html(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
    p := "." + r.URL.Path
    if p == "./" {
        p = "./html/index.html"
    }
    http.ServeFile(w, r, p)
}

func main() {
	http.HandleFunc("/", html)
	http.ListenAndServe(":20800", nil)
}