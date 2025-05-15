package config

import (
	"os"
	"testing"
)

func TesetLoadValidConfig(t *testing.T) {
	config := `[http://google.com, https://github.com]`
	tmpFile, err := os.CreateTemp("", "test_config.json")

	if err != nil {
		t.Fatal("Failed to create temp file:", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test

	urls, err := config.Load(tmpFile.Name())
	if err != nil || len(urls) != 2 {
		t.Errorf("Expected 2 URLs, got %v and error %v", len(urls), err)
	}
}

func TestLoadMissingFile(t *testing.T) {
	_, err := Load("non_existent_file.json")
	if err == nil {
		t.Error("Expected error for missing file, got nil")
	}

}
