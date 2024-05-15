package main

import (
  "internal/routes/routes"
    "fmt"
    "net/http"
)

func main() {
    router := routes.NewRouter()
    fmt.Println("Server started on :8080")
    err := http.ListenAndServe(":8080", router)
    if err!= nil {
        fmt.Println("Error starting server:", err)
    }
}
