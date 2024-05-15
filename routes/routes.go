package routes

import (
	"net/http"

	websockets "github.com/AbhiRam162105/GoBroadcast/internal"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", websockets.HandleConnections)
	return mux
}
