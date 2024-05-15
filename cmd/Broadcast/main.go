package main

import (
	"log"
	"net/http"

	"github.com/AbhiRam162105/GoBroadcast/routes"
)

func main() {
	router := routes.NewRouter()

	log.Println("Server started on :8080")
	err := (http.ListenAndServe(":8080", router))
	if err != nil {
		panic(err)
	}
}
