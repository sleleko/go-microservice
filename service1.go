package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 	   fmt.Fprint(w, "Hello from Service 1!")
    })
    http.ListenAndServe(":8081", nil)
}
