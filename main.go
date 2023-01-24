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
    fmt.Fprintf(w, "Hello, from Docker 🐳\nThis container is made by CI/CD with the image being pushed to both Dockerhub and Azure to build a Web App, all thanks to Github Actions 🔥%s!", r.URL.Path[1:])
}