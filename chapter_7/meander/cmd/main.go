package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/luccafort/building_web_application/chapter_7/meander"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// meander.APIKey = "TODO"
	http.HandleFunc("/journyes", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
