package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/" {
			return
		}

		data, _ := json.Marshal(map[string]time.Time{"currentTime": time.Now()})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":5000", handler)
}
