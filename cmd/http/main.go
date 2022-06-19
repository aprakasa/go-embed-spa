package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aprakasa/go-embed-spa/frontend"
)

func main() {
	http.HandleFunc("/hello.json", handleHello)
	http.HandleFunc("/", handleSPA)
	log.Printf("the server is listening to port %s", os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(map[string]string{
		"message": "hello from the server",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func handleSPA(w http.ResponseWriter, r *http.Request) {
	buildPath := "build"
	f, err := frontend.BuildFs.Open(filepath.Join(buildPath, r.URL.Path))
	if os.IsNotExist(err) {
		index, err := frontend.BuildFs.ReadFile(filepath.Join(buildPath, "index.html"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(index)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	http.FileServer(frontend.BuildHTTPFS()).ServeHTTP(w, r)
}
