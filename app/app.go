package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprinf(w, "Hello DevOps ver2 ")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8081", nil)
}
