package tagfetcher

import (
	"encoding/json"
	"fmt"
	"go-ship/config"
	"go-ship/register"
	"io"
	"net/http"
	"time"
)

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

func FetchLatestTag(registry register.RegistryType, imageName string) (*LatestTag, error) {
	if registry == register.DockerHub {
		return fetchLatestTagFromDockerHub(imageName)
	}
	return nil, fmt.Errorf("unsupported registry type")
}

func FetchWatchedTags(cfg config.ImageConfig) ([]ImageTag, error) {
	var tags ImageTags
	var err error
	if cfg.Registry == register.DockerHub {
		if tags, err = FetchTags(cfg.Name); err != nil {
			return nil, fmt.Errorf("failed to fetch tags for %s: %w", cfg.Name, err)
		}
	}
	if len(cfg.TagPatterns) == 0 {
		// means no tag patterns specified, return the last pushed tag in all tags
		latest := tags.GetLatestPushed()
		if latest == nil {
			return nil, fmt.Errorf("no tags found for %s", cfg.Name)
		}
		return []ImageTag{*latest}, nil
	}

	return nil, fmt.Errorf("unsupported registry type")
}
