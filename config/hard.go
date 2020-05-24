package config

import (
	"fmt"
	"path/filepath"

	"github.com/dvwallin/ago/agotypes"
)

// GetFolders returns all folder paths
func GetFolders() agotypes.Hard {
	return agotypes.Hard{
		PostsFolder:   filepath.Join(".", "posts"),
		SiteFolder:    filepath.Join(".", "site"),
		EntriesFolder: filepath.Join(".", "site", "entries"),
		TagsFolder:    filepath.Join(".", "site", "tags"),
	}
}

// GetStyleFile returns the full path to css file
func GetStyleFile() string {
	cfg := GetCfg()
	return fmt.Sprintf("%s://%sago.css", cfg.Protocol, cfg.Domain)
}
