package tagfetcher

import (
	"path/filepath"
	"time"
)

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

func (it ImageTags) Filter(glob_pattern string) ImageTags {
	var filtered ImageTags
	for _, tag := range it {
		if matchesGlob(tag.Name, glob_pattern) {
			filtered = append(filtered, tag)
		}
	}
	return filtered
}

// Use 3rd party lib. github.com/gobwas/glob
// func matchesGlob(name, pattern string) bool {
// 	g, err := glob.Compile(pattern)
// 	if err != nil {
// 		return false
// 	}
// 	return g.Match(name)
// }

func matchesGlob(name, pattern string) bool {
	matched, err := filepath.Match(pattern, name)
	if err != nil {
		return false // Invalid pattern, no match
	}
	return matched
}
