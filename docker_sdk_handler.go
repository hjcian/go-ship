package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	// "github.com/docker/docker/api/types/filters"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

type RunningContainerInfo struct {
	ContainerID string
	ImageName   string
	Tag         string
	Created     time.Time
}

func getRunningContainerInfo(imageName string) (*RunningContainerInfo, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return nil, err
	}

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
			return &RunningContainerInfo{
				ContainerID: container.ID,
				ImageName:   runningImageName,
				Tag:         tag,
				Created:     time.Unix(container.Created, 0),
			}, nil
		}
	}
	return nil, errors.New("no running container found")
}

func pullImage(imageName string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	resp, err := cli.ImagePull(context.Background(), imageName, image.PullOptions{})
	if err != nil {
		return err
	}
	defer resp.Close()

	return nil
}

func restartContainerWithNewImage(containerID, newImage string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create docker client: %w", err)
	}

	// Get the container's configuration
	containerJSON, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return fmt.Errorf("failed to inspect container: %w", err)
	}

	// Stop the old container
	if err := cli.ContainerStop(context.Background(), containerID, container.StopOptions{}); err != nil {
		return fmt.Errorf("failed to stop container: %w", err)
	}
	// Remove the old container?
	// if err := cli.ContainerRemove(context.Background(), containerID, container.RemoveOptions{}); err != nil {
	// 	return fmt.Errorf("failed to remove container: %w", err)
	// }

	// Create a new container with the same configuration but using the new image
	config := containerJSON.Config
	hostConfig := containerJSON.HostConfig
	networkingConfig := &network.NetworkingConfig{
		EndpointsConfig: containerJSON.NetworkSettings.Networks,
	}

	config.Image = newImage

	newContainer, err := cli.ContainerCreate(context.Background(), config,
		hostConfig,
		networkingConfig,
		nil, containerJSON.Name)
	if err != nil {
		return fmt.Errorf("failed to create new container: %w", err)
	}

	// Start the new container
	if err := cli.ContainerStart(context.Background(), newContainer.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start new container: %w", err)
	}
	return nil
}
