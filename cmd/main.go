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

func respond(w http.ResponseWriter, data any) error {
	return json.NewEncoder(w).Encode(data)
}
