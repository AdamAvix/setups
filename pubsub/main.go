package main

import (
	"net/http"
	"strconv"
	"time"

	"nhooyr.io/websocket"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/subscribe", PublishHandler)
	http.ListenAndServe(":8080", mux)
}

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	client, err := websocket.Accept(w, r, nil)
	handleError(err)
	defer client.Close(websocket.StatusInternalError, "websocket closed unexpectedly")

	for i := 0; ; i++ {
		message := []byte(strconv.Itoa(i))
		client.Write(r.Context(), websocket.MessageText, message)
		time.Sleep(time.Second)
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
