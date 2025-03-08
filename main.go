package main

import (
	"log"
	"time"
)

func mainLoop(image ImageConfig) {
	for {
		log.Printf("Checking for new tags for image: %+v", image)
		tags, err := image.fetchRemoteTags()
		if err != nil {
			log.Printf("Error fetching tags: %v", err)
		} else {
			log.Printf("Fetched tags: %v", tags)
		}

		time.Sleep(3 * time.Second)
	}
}

func main() {
	config, err := LoadConfig("test/test.config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// log.Printf("Loaded config: %+v", config)

	for _, image := range config.Images {
		go mainLoop(image)
	}
	// wait forever
	select {}
}
