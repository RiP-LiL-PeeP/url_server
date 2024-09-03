package main

import (
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, compressURL)
	mux.HandleFunc(`/{id}`, getURL)
	return http.ListenAndServe(`:8080`, mux)
}

func compressURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) //?400
		return
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

}

func getURL(w http.ResponseWriter, r *http.Request) {

}
