package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Yey CICD Jadi")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running on port 4000")
    http.ListenAndServe(":4000", nil)
}
