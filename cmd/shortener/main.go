package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var Shortens = make(map[int]string)

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
	for k, v := range r.Header {
		fmt.Printf("%s: %v\r\n", k, v)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := len(Shortens) + 1
	Shortens[key] = string(body)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprint(key)))
}

func getURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) //?400
		return
	}
	id, err := strconv.Atoi(r.URL.String()[1:])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	url, ok := Shortens[id]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(url)
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusAccepted)
}

/*
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var Shortens = make(map[string]string)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, compressURL)
	mux.HandleFunc(`/`, getURL) // Changed to match all paths, we'll parse them inside the handler
	return http.ListenAndServe(`:8080`, mux)
}

func compressURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	inputURL := string(body)
	if _, err := url.ParseRequestURI(inputURL); err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	hash := sha1.New()
	hash.Write([]byte(inputURL))
	shortURL := hex.EncodeToString(hash.Sum(nil))[:8]

	Shortens[shortURL] = inputURL
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", shortURL)))
}

func getURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/")
	if originalURL, ok := Shortens[id]; ok {
		w.Header().Set("Location", originalURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "URL not found", http.StatusBadRequest)
	}
}

*/