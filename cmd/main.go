package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/taimats/meander"
)

func main() {
	// APIKey := meander.GOOGLE_API_KEY
	ServerPort := os.Getenv("SERVER_PORT")
	http.HandleFunc("/journeys", cors(func(w http.ResponseWriter, r *http.Request) {
		err := respond(w, meander.Journeys)
		if err != nil {
			log.Println(err)
		}
	}))

	http.HandleFunc("/recommendations", cors(func(w http.ResponseWriter, r *http.Request) {
		q := &meander.Query{Journey: strings.Split(r.URL.Query().Get("journey"), "|")}
		q.Lat, _ = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		q.Lng, _ = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
		q.Radius, _ = strconv.Atoi(r.URL.Query().Get("radius"))
		q.CostRangeStr = r.URL.Query().Get("cost")

		places := q.Run()
		respond(w, places)
	}))

	http.ListenAndServe(ServerPort, http.DefaultServeMux)
}

func respond(w http.ResponseWriter, data []any) error {
	publicData := make([]any, len(data))

	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}
