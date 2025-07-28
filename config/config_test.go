package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad_Success(t *testing.T) {
	// Arrange
	tmpDir := t.TempDir()
	yamlContent := `
images:
  - name: "nginx"
    tag_patterns: ["v1.*", "latest"]
    registry: "dockerhub"
  - name: "redis"
    tag_patterns: ["v2.*"]
    registry: "aws"
`
	tmpFile := filepath.Join(tmpDir, "config.yaml")
	if err := os.WriteFile(tmpFile, []byte(yamlContent), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	// Act
	cfg, err := Load(tmpFile)

	// Assert
	assert.NoError(t, err, "expected no error loading config")
	assert.Len(t, cfg.Images, 2, "expected 2 images in config")

	nginx := cfg.Images[0]
	assert.Equal(t, "nginx", nginx.Name, "expected first image name to be 'nginx'")
	assert.Len(t, nginx.TagPatterns, 2, "expected first image to have 2 tag patterns")
	assert.Equal(t, "v1.*", nginx.TagPatterns[0], "expected first tag pattern to be 'v1.*'")
	assert.Equal(t, "latest", nginx.TagPatterns[1], "expected second tag pattern to be 'latest'")
	assert.Equal(t, "dockerhub", nginx.Registry.ToString(), "expected first image registry to be 'dockerhub'")

	redis := cfg.Images[1]
	assert.Equal(t, "redis", redis.Name, "expected second image name to be 'redis'")
	assert.Len(t, redis.TagPatterns, 1, "expected second image to have 1 tag pattern")
	assert.Equal(t, "v2.*", redis.TagPatterns[0], "expected second tag pattern to be 'v2.*'")
	assert.Equal(t, "aws", redis.Registry.ToString(), "expected second image registry to be 'aws'")
}

func TestLoad_FileNotFound(t *testing.T) {
	_, err := Load("nonexistent.yaml")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "bad.yaml")
	badYAML := "images: [name: nginx, tag_pattern: v1.*" // malformed YAML
	if err := os.WriteFile(tmpFile, []byte(badYAML), 0644); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	_, err := Load(tmpFile)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}
