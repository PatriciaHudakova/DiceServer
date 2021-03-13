package main

import (
	"fmt"
	dice "github.com/PatriciaHudakova/DiceLibrary"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/GET/v1/roll", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Your rolled value is: %v", dice.Roll())
	})

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
