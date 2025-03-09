package main

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	// "github.com/docker/docker/api/types/filters"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type RunningContainerInfo struct {
	ImageName string
	Tag       string
	Created   time.Time
}

func getRunningContainersInfo() ([]container.Summary, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func getRunningContainerTag(imageName string) (*RunningContainerInfo, error) {
	containers, err := getRunningContainersInfo()
	if err != nil {
		return nil, err
	}
	for _, container := range containers {
		runningImageName := container.Image
		tag := "latest" // If no tag is specified, default to latest.

		if strings.Contains(container.Image, ":") {
			parts := strings.SplitN(container.Image, ":", 2)
			runningImageName = parts[0]
			tag = parts[1]
		}
		if runningImageName == imageName && container.State == "running" {
			// printJSON("Container", container)
			return &RunningContainerInfo{
				ImageName: runningImageName,
				Tag:       tag,
				Created:   time.Unix(container.Created, 0),
			}, nil
		}
	}
	return nil, errors.New("no running container found")
}

func getLocalTags(imageName string) ([]string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		return nil, err
	}
	log.Printf("Found %d images", len(images))
	for _, image := range images {
		isContains := false
		for _, repoTag := range image.RepoTags {
			if strings.Split(repoTag, ":")[0] == imageName {
				isContains = true
			}
		}
		if !isContains {
			continue
		}

		printJSON("Image", image)
	}
	// log.Printf("Found %v images", images)
	var tags []string
	// for _, img := range images {
	// 	for _, tag := range img.RepoTags {
	// 		if tag == imageName || tag == imageName+":latest" {
	// 			tags = append(tags, tag)
	// 		}
	// 	}
	// }

	return tags, nil
}
