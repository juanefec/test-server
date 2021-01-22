package server

import (
	"fmt"
	"net/http"
)

// StartServer serves a server that serves.
func StartServer(port int) {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
}
