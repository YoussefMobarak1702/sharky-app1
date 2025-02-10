package main;
import (
    "fmt"
    "log"
    "net/http"
)
func main() {
    http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "There's no such thing as a bad day when you're fishing.")
    })
    fmt.Printf("Server running (port=3001), route: http://localhost:3001/helloworld\n")
    if err := http.ListenAndServe(":3001", nil); err != nil {
        log.Fatal(err)
    }
}
