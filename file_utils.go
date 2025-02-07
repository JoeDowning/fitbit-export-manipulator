package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func extractJSONFiles(path string) (FilesMap, error) {
	files, err := getJSONFiles(path)
	if err != nil {
		return FilesMap{}, fmt.Errorf("failed to get JSON files: %w", err)
	}

	filesMap := make(map[string]Entries)
	for _, file := range files {
		entries, err := extractEntries(path, file)
		if err != nil {
			return FilesMap{}, fmt.Errorf("failed to extract entries: %w", err)
		}

		filesMap[file] = entries
	}

	return filesMap, nil
}

func extractEntries(path, file string) (Entries, error) {
	filePath := fmt.Sprintf("%s/%s", path, file)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var entries Entries
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return entries, nil
}

func getJSONFiles(path string) ([]string, error) {
	entries, err := getDirectoryEntries(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get directory entries: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if isJSONFile(entry.Name()) {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

func getDirectoryEntries(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	return entries, nil
}

func isJSONFile(name string) bool {
	return strings.HasSuffix(name, ".json")
}
