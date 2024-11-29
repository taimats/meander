package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/taimats/meander"
)

func main() {
	//meader.APIKey = "tmp"
	ServerPort := os.Getenv("SERVER_PORT")
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		err := respond(w, meander.Journeys)
		if err != nil {
			log.Println(err)
		}
	})

	http.ListenAndServe(ServerPort, http.DefaultServeMux)

}

func respond(w http.ResponseWriter, data []any) error {
	publicData := make([]any, len(data))

	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
