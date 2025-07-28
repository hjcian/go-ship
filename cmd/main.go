package main

import (
	"encoding/json"
	"go-ship/tagfetcher"
	"log"
	"os"
)

func main() {
	resp, _ := tagfetcher.FetchTags("redis")

	// Write resp to json file
	if resp == nil {
		return
	}
	// Convert to JSON
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	// Write to file
	err = os.WriteFile("redis_tags.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	log.Println("Successfully wrote tags to redis_tags.json")
}
