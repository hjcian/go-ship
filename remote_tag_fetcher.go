package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Struct to parse the JSON response

type RepoTagsResp struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Creator int `json:"creator"`
		ID      int `json:"id"`
		Images  []struct {
			Architecture string      `json:"architecture"`
			Features     string      `json:"features"`
			Variant      interface{} `json:"variant"`
			Digest       string      `json:"digest"`
			Os           string      `json:"os"`
			OsFeatures   string      `json:"os_features"`
			OsVersion    interface{} `json:"os_version"`
			Size         int         `json:"size"`
			Status       string      `json:"status"`
			LastPulled   time.Time   `json:"last_pulled"`
			LastPushed   time.Time   `json:"last_pushed"`
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
		Digest              string    `json:"digest"`
	} `json:"results"`
}

func (i *ImageConfig) fetchRemoteTags() (*RepoTagsResp, error) {
	if i.Registry == DockerHub {

		url := fmt.Sprintf("https://registry.hub.docker.com/v2/namespaces/library/repositories/%s/tags?page_size=1", i.Name)
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
		// print the response
		fmt.Println(string(body))

		var tagReps RepoTagsResp
		if err := json.Unmarshal(body, &tagReps); err != nil {
			return nil, err
		}

		return &tagReps, nil
	}

	return nil, fmt.Errorf("unsupported registry type")
}
