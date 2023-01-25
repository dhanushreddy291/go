package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, from Docker🐳\nThis container image is made by Github Actions and deployed to Cloud Run via Cloud Build%s!", r.URL.Path[1:])
}