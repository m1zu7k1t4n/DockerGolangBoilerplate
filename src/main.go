package main
import (
    "fmt"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/api", handler_api)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Your url path is %s", r.URL.Path[1:])
}

func handler_api(w http.ResponseWriter, r *http.Request) {
    keys, ok := r.URL.Query()["key"]
    if !ok || len(keys[0]) < 1 {
        fmt.Fprintf(w, "api parameter: None")
        return
    }

    fmt.Fprintf(w, "api parameter: %s", keys)
}
