package tagfetcher

import "time"

type ImageTag struct {
	Name          string    `json:"name"`            // The tag name
	TagLastPushed time.Time `json:"tag_last_pushed"` // NOTE: the last pushed time
	Digest        string    `json:"digest"`          // NOTE: image ID
}

type ImageTags []ImageTag

func (it ImageTags) Len() int {
	return len(it)
}

func (it ImageTags) GetLatestPushed() *ImageTag {
	if len(it) == 0 {
		return nil
	}
	latest := it[0]
	for _, tag := range it {
		if tag.TagLastPushed.After(latest.TagLastPushed) {
			latest = tag
		}
	}
	return &latest
}
