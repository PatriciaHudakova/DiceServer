package main

import (
	"encoding/json"
	"log"
	"net/http"

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

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
