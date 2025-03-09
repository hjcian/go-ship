package main

import (
	"encoding/json"
	"log"
)

func printJSON(message string, data any) {
	imageJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling image to JSON: %v", err)
	}
	log.Printf("%s: %s", message, imageJSON)
}
