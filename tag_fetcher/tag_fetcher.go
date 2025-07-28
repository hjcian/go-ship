package tagfetcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RegistryType int

const (
	DockerHub RegistryType = iota
	AWS
)

func (r *RegistryType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	switch s {
	case "dockerhub":
		*r = DockerHub
	case "aws":
		*r = AWS
	default:
		return errors.New("invalid registry type")
	}

	return nil
}

type LatestTag struct {
	TagName    string
	LastPushed time.Time
}

func fetchLatestTagFromDockerHub(imageName string) (*LatestTag, error) {
	url := fmt.Sprintf("https://registry.hub.docker.com/v2/namespaces/library/repositories/%s/tags?page_size=1", imageName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch tags: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tagReps RepoTagsResp
	if err := json.Unmarshal(body, &tagReps); err != nil {
		return nil, err
	}

	return &LatestTag{
		TagName:    tagReps.Results[0].Name,
		LastPushed: tagReps.Results[0].TagLastPushed,
	}, nil
}

func FetchLatestTag(registry RegistryType, imageName string) (*LatestTag, error) {
	if registry == DockerHub {
		return fetchLatestTagFromDockerHub(imageName)
	}
	return nil, fmt.Errorf("unsupported registry type")
}

type RepoTagsResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Creator int `json:"creator"`
		ID      int `json:"id"`
		Images  []struct {
			Architecture string    `json:"architecture"`
			Features     string    `json:"features"`
			Variant      any       `json:"variant"`
			Digest       string    `json:"digest"`
			Os           string    `json:"os"`
			OsFeatures   string    `json:"os_features"`
			OsVersion    any       `json:"os_version"`
			Size         int       `json:"size"`
			Status       string    `json:"status"`
			LastPulled   time.Time `json:"last_pulled"`
			LastPushed   time.Time `json:"last_pushed"`
		} `json:"images"`
		LastUpdated         time.Time `json:"last_updated"`
		LastUpdater         int       `json:"last_updater"`
		LastUpdaterUsername string    `json:"last_updater_username"`
		Name                string    `json:"name"` // NOTE: the tag name
		Repository          int       `json:"repository"`
		FullSize            int       `json:"full_size"`
		V2                  bool      `json:"v2"`
		TagStatus           string    `json:"tag_status"`
		TagLastPulled       time.Time `json:"tag_last_pulled"`
		TagLastPushed       time.Time `json:"tag_last_pushed"` // NOTE: the last pushed time
		MediaType           string    `json:"media_type"`
		ContentType         string    `json:"content_type"`
		Digest              string    `json:"digest"` // NOTE: image ID
	} `json:"results"`
}
