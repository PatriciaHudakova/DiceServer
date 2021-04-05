package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	dice "github.com/PatriciaHudakova/DiceLibrary"
)

type rollingDice struct {
	Roll int `json:"value"`
}

func main() {
	http.HandleFunc("/v1/roll", func(w http.ResponseWriter, r *http.Request) {
		roll := rollingDice{
			Roll: dice.Roll(),
		}

		log.Println("the dice has been rolled")

		if err := json.NewEncoder(w).Encode(roll); err != nil {
			log.Printf("unable to marshall json: %s", err)
		}
	})

    // log messages
	log.Printf("Listening on localhost %s", getPort())

	log.Fatal(http.ListenAndServe(getPort(), nil))
}

// getPort gets the Port from env
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
