package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 	   fmt.Fprint(w, "Hello from Service 2!")
    })
    http.ListenAndServe(":8082", nil)
}
