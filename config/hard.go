package config

import (
	"path/filepath"
)

// Hard is for hard coded config values
type Hard struct {
	PostsFolder   string
	SiteFolder    string
	EntriesFolder string
	TagsFolder    string
}

// GetFolders returns all folder paths
func GetFolders() Hard {
	return Hard{
		PostsFolder:   filepath.Join(".", "posts"),
		SiteFolder:    filepath.Join(".", "site"),
		EntriesFolder: filepath.Join(".", "site", "entries"),
		TagsFolder:    filepath.Join(".", "site", "tags"),
	}
}
