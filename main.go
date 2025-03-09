package main

import (
	"log"
	"time"
)

func mainLoop(image ImageConfig) {
	first := true
start:
	for {
		if !first {
			time.Sleep(3 * time.Second)
		} else {
			first = false
		}

		log.Printf("Checking for new tags for image: %+v", image)

		tag, err := getRunningContainerTag(image.Name)
		if err != nil {
			log.Printf("Error fetching local tag: %v", err)
			goto start
		}
		log.Println("Found tag of running container:", tag)

		latestTag, err := fetchLatestTag(image.Registry, image.Name)
		if err != nil {
			log.Printf("Error fetching tags: %v", err)
		} else {
			log.Printf("Fetched tags: %v", latestTag)
		}
	}
}

func main() {
	config, err := LoadConfig("test/test.config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	for _, image := range config.Images {
		go mainLoop(image)
	}
	// wait forever
	select {}
}
