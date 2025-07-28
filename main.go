package main

import (
	tagfetcher "go-ship/tag_fetcher"
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

		runningContainerInfo, err := getRunningContainerInfo(image.Name)
		if err != nil {
			log.Printf("Error fetching local tag: %v", err)
			goto start
		}
		log.Println("Found tag of running container:", runningContainerInfo)

		latestTag, err := tagfetcher.FetchLatestTag(image.Registry, image.Name)
		if err != nil {
			log.Printf("Error fetching tags: %v", err)
			goto start
		}
		log.Printf("Fetched tags: %v", latestTag)

		if runningContainerInfo.Tag == latestTag.TagName {
			log.Println("Tags are the same, no need to update")
			goto start
		}

		log.Printf("Pulling image: %s:%s", image.Name, latestTag.TagName)
		newImage := image.Name + ":" + latestTag.TagName
		if err := pullImage(newImage); err != nil {
			log.Printf("Error pulling image: %v", err)
			goto start
		}
		log.Println("Pulled image successfully")

		log.Println("Restarting container...")
		if err := restartContainerWithNewImage(runningContainerInfo.ContainerID, newImage); err != nil {
			log.Printf("Error restarting container: %v", err)
			goto start
		}
		log.Println("Restarted container successfully")
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
