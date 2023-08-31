package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", HelloServer)

    server_port := ":8080"

    http.ListenAndServe(server_port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, os.Getenv("GREETING") + " world - We like: Yellow . ")
}
